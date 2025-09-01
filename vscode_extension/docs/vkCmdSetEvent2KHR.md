# vkCmdSetEvent2(3) Manual Page

## Name

vkCmdSetEvent2 - Set an event object to signaled state



## [](#_c_specification)C Specification

To signal an event from a device, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdSetEvent2(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    const VkDependencyInfo*                     pDependencyInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
void vkCmdSetEvent2KHR(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    const VkDependencyInfo*                     pDependencyInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `event` is the event that will be signaled.
- `pDependencyInfo` is a pointer to a [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structure defining the first scopes of this operation.

## [](#_description)Description

When [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) is submitted to a queue, it defines the first half of memory dependencies defined by `pDependencyInfo`, as well as an event signal operation which sets the event to the signaled state. A memory dependency is defined between the event signal operation and commands that occur earlier in submission order.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) and [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) are defined by the union of all the memory dependencies defined by `pDependencyInfo`, and are applied to all operations that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). [Queue family ownership transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) and [image layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions) defined by `pDependencyInfo` are also included in the first scopes.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the event signal operation, and any [queue family ownership transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) and [image layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions) defined by `pDependencyInfo`.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) includes only [queue family ownership transfers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) and [image layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).

If `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR` is not set in `pDependencyInfo->dependencyFlags`, future [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) commands rely on all values of each element in `pDependencyInfo` matching exactly with those used to signal the corresponding event. If `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR` is set, `vkCmdSetEvent2` **must** only include the [source stage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-masks) of the first synchronization scope in `pDependencyInfo->pMemoryBarriers`\[0].`srcStageMask`. [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) **must** not be used to wait on the result of a signal operation defined by `vkCmdSetEvent2`.

Note

The extra information provided by [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) compared to [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html) allows implementations to more efficiently schedule the operations required to satisfy the requested dependencies. With [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html), the full dependency information is not known until [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) is recorded, forcing implementations to insert the required operations at that point and not before.

If `event` is already in the signaled state when [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) is executed on the device, then [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent2.html) has no effect, no event signal operation occurs, and no dependency is generated.

Valid Usage

- [](#VUID-vkCmdSetEvent2-synchronization2-03824)VUID-vkCmdSetEvent2-synchronization2-03824  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdSetEvent2-dependencyFlags-03825)VUID-vkCmdSetEvent2-dependencyFlags-03825  
  The `dependencyFlags` member of `pDependencyInfo` **must** be `0` or `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`
- [](#VUID-vkCmdSetEvent2-srcStageMask-09391)VUID-vkCmdSetEvent2-srcStageMask-09391  
  The `srcStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfo` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-vkCmdSetEvent2-dstStageMask-09392)VUID-vkCmdSetEvent2-dstStageMask-09392  
  The `dstStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfo` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`
- [](#VUID-vkCmdSetEvent2-commandBuffer-03826)VUID-vkCmdSetEvent2-commandBuffer-03826  
  The current device mask of `commandBuffer` **must** include exactly one physical device
- [](#VUID-vkCmdSetEvent2-srcStageMask-03827)VUID-vkCmdSetEvent2-srcStageMask-03827  
  The `srcStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdSetEvent2-dstStageMask-03828)VUID-vkCmdSetEvent2-dstStageMask-03828  
  The `dstStageMask` member of any element of the `pMemoryBarriers`, `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of `pDependencyInfo` **must** only include pipeline stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdSetEvent2-dependencyFlags-10785)VUID-vkCmdSetEvent2-dependencyFlags-10785  
  If the `dependencyFlags` member of `pDependencyInfo` includes `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, the `bufferMemoryBarrierCount` and `imageMemoryBarrierCount` members of `pDependencyInfo` **must** be `0`
- [](#VUID-vkCmdSetEvent2-dependencyFlags-10786)VUID-vkCmdSetEvent2-dependencyFlags-10786  
  If the `dependencyFlags` member of `pDependencyInfo` includes `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, the `memoryBarrierCount` member of `pDependencyInfo` **must** be `1`
- [](#VUID-vkCmdSetEvent2-dependencyFlags-10787)VUID-vkCmdSetEvent2-dependencyFlags-10787  
  If the `dependencyFlags` member of `pDependencyInfo` includes `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`, the `srcAccessMask`, `dstStageMask`, and `dstAccessMask` members of `pDependencyInfo->pMemoryBarriers`\[0] **must** be `0`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetEvent2-commandBuffer-parameter)VUID-vkCmdSetEvent2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetEvent2-event-parameter)VUID-vkCmdSetEvent2-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkCmdSetEvent2-pDependencyInfo-parameter)VUID-vkCmdSetEvent2-pDependencyInfo-parameter  
  `pDependencyInfo` **must** be a valid pointer to a valid [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html) structure
- [](#VUID-vkCmdSetEvent2-commandBuffer-recording)VUID-vkCmdSetEvent2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetEvent2-commandBuffer-cmdpool)VUID-vkCmdSetEvent2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, or encode operations
- [](#VUID-vkCmdSetEvent2-renderpass)VUID-vkCmdSetEvent2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdSetEvent2-commonparent)VUID-vkCmdSetEvent2-commonparent  
  Both of `commandBuffer`, and `event` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Both

Graphics  
Compute  
Decode  
Encode

Synchronization

Conditional Rendering

vkCmdSetEvent2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetEvent2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0