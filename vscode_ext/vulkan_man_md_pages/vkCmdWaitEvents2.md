# vkCmdWaitEvents2(3) Manual Page

## Name

vkCmdWaitEvents2 - Wait for one or more events



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait for one or more events to enter the signaled state on a device,
call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdWaitEvents2(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    eventCount,
    const VkEvent*                              pEvents,
    const VkDependencyInfo*                     pDependencyInfos);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
void vkCmdWaitEvents2KHR(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    eventCount,
    const VkEvent*                              pEvents,
    const VkDependencyInfo*                     pDependencyInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `eventCount` is the length of the `pEvents` array.

- `pEvents` is a pointer to an array of `eventCount` events to wait on.

- `pDependencyInfos` is a pointer to an array of `eventCount`
  [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html) structures, defining the
  second <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">synchronization scope</a>.

## <a href="#_description" class="anchor"></a>Description

When `vkCmdWaitEvents2` is submitted to a queue, it inserts memory
dependencies according to the elements of `pDependencyInfos` and each
corresponding element of `pEvents`. `vkCmdWaitEvents2` **must** not be
used to wait on event signal operations occurring on other queues, or
signal operations executed by [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html).

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> of each memory
dependency defined by any element i of `pDependencyInfos` are applied to
operations that occurred earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> than the last event
signal operation on element i of `pEvents`.

Signal operations for an event at index i are only included if:

- The event was signaled by a [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html)
  command that occurred earlier in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
  target="_blank" rel="noopener">submission order</a> with a
  `dependencyInfo` parameter exactly equal to the element of
  `pDependencyInfos` at index i ; or

- The event was created without `VK_EVENT_CREATE_DEVICE_ONLY_BIT`, and
  the first <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">synchronization scope</a> defined by
  the element of `pDependencyInfos` at index i only includes host
  operations (`VK_PIPELINE_STAGE_2_HOST_BIT`).

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> of each memory
dependency defined by any element i of `pDependencyInfos` are applied to
operations that occurred later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> than
`vkCmdWaitEvents2`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html">vkCmdWaitEvents2</a> is used with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html">vkCmdSetEvent2</a> to define a memory
dependency between two sets of action commands, roughly in the same way
as pipeline barriers, but split into two commands such that work between
the two <strong>may</strong> execute unhindered.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications should be careful to avoid race conditions when using
events. There is no direct ordering guarantee between
<code>vkCmdSetEvent2</code> and <a
href="vkCmdResetEvent2.html">vkCmdResetEvent2</a>, <a
href="vkCmdResetEvent.html">vkCmdResetEvent</a>, or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html">vkCmdSetEvent</a>. Another execution
dependency (e.g. a pipeline barrier or semaphore with
<code>VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT</code>) is needed to prevent
such a race condition.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdWaitEvents2-synchronization2-03836"
  id="VUID-vkCmdWaitEvents2-synchronization2-03836"></a>
  VUID-vkCmdWaitEvents2-synchronization2-03836  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdWaitEvents2-pEvents-03837"
  id="VUID-vkCmdWaitEvents2-pEvents-03837"></a>
  VUID-vkCmdWaitEvents2-pEvents-03837  
  Members of `pEvents` **must** not have been signaled by
  [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html)

- <a href="#VUID-vkCmdWaitEvents2-pEvents-03838"
  id="VUID-vkCmdWaitEvents2-pEvents-03838"></a>
  VUID-vkCmdWaitEvents2-pEvents-03838  
  For any element i of `pEvents`, if that event is signaled by
  [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html), that command’s `dependencyInfo`
  parameter **must** be exactly equal to the ith element of
  `pDependencyInfos`

- <a href="#VUID-vkCmdWaitEvents2-pEvents-03839"
  id="VUID-vkCmdWaitEvents2-pEvents-03839"></a>
  VUID-vkCmdWaitEvents2-pEvents-03839  
  For any element i of `pEvents`, if that event is signaled by
  [vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html), barriers in the ith element of
  `pDependencyInfos` **must** include only host operations in their
  first <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
  target="_blank" rel="noopener">synchronization scope</a>

- <a href="#VUID-vkCmdWaitEvents2-pEvents-03840"
  id="VUID-vkCmdWaitEvents2-pEvents-03840"></a>
  VUID-vkCmdWaitEvents2-pEvents-03840  
  For any element i of `pEvents`, if barriers in the ith element of
  `pDependencyInfos` include only host operations, the ith element of
  `pEvents` **must** be signaled before
  [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html) is executed

- <a href="#VUID-vkCmdWaitEvents2-pEvents-03841"
  id="VUID-vkCmdWaitEvents2-pEvents-03841"></a>
  VUID-vkCmdWaitEvents2-pEvents-03841  
  For any element i of `pEvents`, if barriers in the ith element of
  `pDependencyInfos` do not include host operations, the ith element of
  `pEvents` **must** be signaled by a corresponding
  [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html) that occurred earlier in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
  target="_blank" rel="noopener">submission order</a>

- <a href="#VUID-vkCmdWaitEvents2-srcStageMask-03842"
  id="VUID-vkCmdWaitEvents2-srcStageMask-03842"></a>
  VUID-vkCmdWaitEvents2-srcStageMask-03842  
  The `srcStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfos` **must** either include only pipeline stages valid
  for the queue family that was used to create the command pool that
  `commandBuffer` was allocated from

- <a href="#VUID-vkCmdWaitEvents2-dstStageMask-03843"
  id="VUID-vkCmdWaitEvents2-dstStageMask-03843"></a>
  VUID-vkCmdWaitEvents2-dstStageMask-03843  
  The `dstStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfos` **must** only include pipeline stages valid for the
  queue family that was used to create the command pool that
  `commandBuffer` was allocated from

- <a href="#VUID-vkCmdWaitEvents2-dependencyFlags-03844"
  id="VUID-vkCmdWaitEvents2-dependencyFlags-03844"></a>
  VUID-vkCmdWaitEvents2-dependencyFlags-03844  
  If `vkCmdWaitEvents2` is being called inside a render pass instance,
  the `srcStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfos` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-vkCmdWaitEvents2-commandBuffer-03846"
  id="VUID-vkCmdWaitEvents2-commandBuffer-03846"></a>
  VUID-vkCmdWaitEvents2-commandBuffer-03846  
  `commandBuffer`’s current device mask **must** include exactly one
  physical device

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWaitEvents2-commandBuffer-parameter"
  id="VUID-vkCmdWaitEvents2-commandBuffer-parameter"></a>
  VUID-vkCmdWaitEvents2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWaitEvents2-pEvents-parameter"
  id="VUID-vkCmdWaitEvents2-pEvents-parameter"></a>
  VUID-vkCmdWaitEvents2-pEvents-parameter  
  `pEvents` **must** be a valid pointer to an array of `eventCount`
  valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handles

- <a href="#VUID-vkCmdWaitEvents2-pDependencyInfos-parameter"
  id="VUID-vkCmdWaitEvents2-pDependencyInfos-parameter"></a>
  VUID-vkCmdWaitEvents2-pDependencyInfos-parameter  
  `pDependencyInfos` **must** be a valid pointer to an array of
  `eventCount` valid [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html)
  structures

- <a href="#VUID-vkCmdWaitEvents2-commandBuffer-recording"
  id="VUID-vkCmdWaitEvents2-commandBuffer-recording"></a>
  VUID-vkCmdWaitEvents2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWaitEvents2-commandBuffer-cmdpool"
  id="VUID-vkCmdWaitEvents2-commandBuffer-cmdpool"></a>
  VUID-vkCmdWaitEvents2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdWaitEvents2-eventCount-arraylength"
  id="VUID-vkCmdWaitEvents2-eventCount-arraylength"></a>
  VUID-vkCmdWaitEvents2-eventCount-arraylength  
  `eventCount` **must** be greater than `0`

- <a href="#VUID-vkCmdWaitEvents2-commonparent"
  id="VUID-vkCmdWaitEvents2-commonparent"></a>
  VUID-vkCmdWaitEvents2-commonparent  
  Both of `commandBuffer`, and the elements of `pEvents` **must** have
  been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td
class="tableblock halign-left valign-top"><p>Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWaitEvents2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
