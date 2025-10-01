# vkCmdWaitEvents2(3) Manual Page

## Name

vkCmdWaitEvents2 - Wait for one or more events



## [](#_c_specification)C Specification

To wait for one or more events to enter the signaled state on a device, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdWaitEvents2(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    eventCount,
    const VkEvent*                              pEvents,
    const VkDependencyInfo*                     pDependencyInfos);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
void vkCmdWaitEvents2KHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    eventCount,
    const VkEvent*                              pEvents,
    const VkDependencyInfo*                     pDependencyInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `eventCount` is the length of the `pEvents` array.
- `pEvents` is a pointer to an array of `eventCount` events to wait on.
- `pDependencyInfos` is a pointer to an array of `eventCount` [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structures, defining the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes).

## [](#_description)Description

When `vkCmdWaitEvents2` is submitted to a queue, it inserts memory dependencies according to the elements of `pDependencyInfos` and each corresponding element of `pEvents`. `vkCmdWaitEvents2` **must** not be used to wait on event signal operations occurring on other queues, or signal operations executed by [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html).

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) of each memory dependency defined by any element i of `pDependencyInfos` are applied to operations that occurred earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) than the last event signal operation on element i of `pEvents`.

Signal operations for an event at index i are only included if:

- The event was signaled by a [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) command that occurred earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) with a `dependencyInfo` parameter exactly equal to the element of `pDependencyInfos` at index i ; or
- The event was created without `VK_EVENT_CREATE_DEVICE_ONLY_BIT`, and the first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) defined by the element of `pDependencyInfos` at index i only includes host operations (`VK_PIPELINE_STAGE_2_HOST_BIT`).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) of each memory dependency defined by any element i of `pDependencyInfos` are applied to operations that occurred later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) than `vkCmdWaitEvents2`.

Note

[vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) is used with [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) to define a memory dependency between two sets of action commands, roughly in the same way as pipeline barriers, but split into two commands such that work between the two **may** execute unhindered.

Note

Applications should be careful to avoid race conditions when using events. There is no direct ordering guarantee between `vkCmdSetEvent2` and [vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent2.html), [vkCmdResetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent.html), or [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html). Another execution dependency (e.g. a pipeline barrier or semaphore with `VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT`) is needed to prevent such a race condition.

Valid Usage

- [](#VUID-vkCmdWaitEvents2-synchronization2-03836)VUID-vkCmdWaitEvents2-synchronization2-03836  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdWaitEvents2-pEvents-03837)VUID-vkCmdWaitEvents2-pEvents-03837  
  Members of `pEvents` **must** not have been signaled by [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html)
- [](#VUID-vkCmdWaitEvents2-pEvents-10788)VUID-vkCmdWaitEvents2-pEvents-10788  
  For any element i of `pEvents`, if the `dependencyFlags` member of the ith element of `pDependencyInfos` does not include `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, and if that event is signaled by [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html), that command’s `dependencyInfo` parameter **must** be exactly equal to the ith element of `pDependencyInfos`
- [](#VUID-vkCmdWaitEvents2-pEvents-10789)VUID-vkCmdWaitEvents2-pEvents-10789  
  For any element i of `pEvents`, if the `dependencyFlags` member of the ith element of `pDependencyInfos` includes `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, that event **must** be signaled by [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) with `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`
- [](#VUID-vkCmdWaitEvents2-pEvents-10790)VUID-vkCmdWaitEvents2-pEvents-10790  
  For any element i of `pEvents`, if the `dependencyFlags` member of the ith element of `pDependencyInfos` includes `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, the union of `srcStageMask` members of all elements of `pMemoryBarriers`, `pBufferMemoryBarriers`, and `pImageMemoryBarriers` of the ith element of `pDependencyInfos` **must** equal `pDependencyInfos->pMemoryBarriers`\[0].`srcStageMask` in the [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) command
- [](#VUID-vkCmdWaitEvents2-pEvents-03839)VUID-vkCmdWaitEvents2-pEvents-03839  
  For any element i of `pEvents`, if that event is signaled by [vkSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetEvent.html), barriers in the ith element of `pDependencyInfos` **must** include only host operations in their first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes)
- [](#VUID-vkCmdWaitEvents2-pEvents-03840)VUID-vkCmdWaitEvents2-pEvents-03840  
  For any element i of `pEvents`, if barriers in the ith element of `pDependencyInfos` include only host operations, the ith element of `pEvents` **must** be signaled before [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) is executed
- [](#VUID-vkCmdWaitEvents2-pEvents-03841)VUID-vkCmdWaitEvents2-pEvents-03841  
  For any element i of `pEvents`, if barriers in the ith element of `pDependencyInfos` do not include host operations, the ith element of `pEvents` **must** be signaled by a corresponding [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) that occurred earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order)
- [](#VUID-vkCmdWaitEvents2-srcStageMask-03842)VUID-vkCmdWaitEvents2-srcStageMask-03842  
  The `srcStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfos` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWaitEvents2-dstStageMask-03843)VUID-vkCmdWaitEvents2-dstStageMask-03843  
  The `dstStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfos` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWaitEvents2-dependencyFlags-10394)VUID-vkCmdWaitEvents2-dependencyFlags-10394  
  The `dependencyFlags` member of any element of `pDependencyInfo` **must** not include any of the following bits:
  
  - `VK_DEPENDENCY_BY_REGION_BIT`
  - `VK_DEPENDENCY_DEVICE_GROUP_BIT`
  - `VK_DEPENDENCY_VIEW_LOCAL_BIT`
  - `VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT`
- [](#VUID-vkCmdWaitEvents2-maintenance8-10205)VUID-vkCmdWaitEvents2-maintenance8-10205  
  If the [`maintenance8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance8) feature is not enabled, the `dependencyFlags` members of any element of `pDependencyInfos` **must** not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`
- [](#VUID-vkCmdWaitEvents2-dependencyFlags-03844)VUID-vkCmdWaitEvents2-dependencyFlags-03844  
  If this command is called inside a render pass instance, the `srcStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfos` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-vkCmdWaitEvents2-commandBuffer-03846)VUID-vkCmdWaitEvents2-commandBuffer-03846  
  `commandBuffer`’s current device mask **must** include exactly one physical device
- [](#VUID-vkCmdWaitEvents2-None-10654)VUID-vkCmdWaitEvents2-None-10654  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdWaitEvents2-commandBuffer-parameter)VUID-vkCmdWaitEvents2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWaitEvents2-pEvents-parameter)VUID-vkCmdWaitEvents2-pEvents-parameter  
  `pEvents` **must** be a valid pointer to an array of `eventCount` valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handles
- [](#VUID-vkCmdWaitEvents2-pDependencyInfos-parameter)VUID-vkCmdWaitEvents2-pDependencyInfos-parameter  
  `pDependencyInfos` **must** be a valid pointer to an array of `eventCount` valid [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structures
- [](#VUID-vkCmdWaitEvents2-commandBuffer-recording)VUID-vkCmdWaitEvents2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWaitEvents2-commandBuffer-cmdpool)VUID-vkCmdWaitEvents2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations
- [](#VUID-vkCmdWaitEvents2-eventCount-arraylength)VUID-vkCmdWaitEvents2-eventCount-arraylength  
  `eventCount` **must** be greater than `0`
- [](#VUID-vkCmdWaitEvents2-commonparent)VUID-vkCmdWaitEvents2-commonparent  
  Both of `commandBuffer`, and the elements of `pEvents` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Synchronization

Conditional Rendering

vkCmdWaitEvents2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWaitEvents2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0