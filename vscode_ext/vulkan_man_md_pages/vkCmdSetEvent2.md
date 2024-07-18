# vkCmdSetEvent2(3) Manual Page

## Name

vkCmdSetEvent2 - Set an event object to signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To signal an event from a device, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdSetEvent2(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    const VkDependencyInfo*                     pDependencyInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
void vkCmdSetEvent2KHR(
    VkCommandBuffer                             commandBuffer,
    VkEvent                                     event,
    const VkDependencyInfo*                     pDependencyInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `event` is the event that will be signaled.

- `pDependencyInfo` is a pointer to a
  [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html) structure defining the first
  scopes of this operation.

## <a href="#_description" class="anchor"></a>Description

When [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html) is submitted to a queue, it
defines the first half of memory dependencies defined by
`pDependencyInfo`, as well as an event signal operation which sets the
event to the signaled state. A memory dependency is defined between the
event signal operation and commands that occur earlier in submission
order.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> are defined by the union
of all the memory dependencies defined by `pDependencyInfo`, and are
applied to all operations that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">Queue family ownership transfers</a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a> defined by
`pDependencyInfo` are also included in the first scopes.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes only
the event signal operation, and any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfers</a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a> defined by
`pDependencyInfo`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> includes only <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfers</a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a>.

Future [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html) commands rely on all
values of each element in `pDependencyInfo` matching exactly with those
used to signal the corresponding event.
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) **must** not be used to wait on
the result of a signal operation defined by `vkCmdSetEvent2`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The extra information provided by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html">vkCmdSetEvent2</a> compared to <a
href="vkCmdSetEvent.html">vkCmdSetEvent</a> allows implementations to
more efficiently schedule the operations required to satisfy the
requested dependencies. With <a
href="vkCmdSetEvent.html">vkCmdSetEvent</a>, the full dependency
information is not known until <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html">vkCmdWaitEvents</a> is recorded, forcing
implementations to insert the required operations at that point and not
before.</p></td>
</tr>
</tbody>
</table>

If `event` is already in the signaled state when
[vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html) is executed on the device, then
[vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html) has no effect, no event signal
operation occurs, and no dependency is generated.

Valid Usage

- <a href="#VUID-vkCmdSetEvent2-synchronization2-03824"
  id="VUID-vkCmdSetEvent2-synchronization2-03824"></a>
  VUID-vkCmdSetEvent2-synchronization2-03824  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdSetEvent2-dependencyFlags-03825"
  id="VUID-vkCmdSetEvent2-dependencyFlags-03825"></a>
  VUID-vkCmdSetEvent2-dependencyFlags-03825  
  The `dependencyFlags` member of `pDependencyInfo` **must** be `0`

- <a href="#VUID-vkCmdSetEvent2-srcStageMask-09391"
  id="VUID-vkCmdSetEvent2-srcStageMask-09391"></a>
  VUID-vkCmdSetEvent2-srcStageMask-09391  
  The `srcStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-vkCmdSetEvent2-dstStageMask-09392"
  id="VUID-vkCmdSetEvent2-dstStageMask-09392"></a>
  VUID-vkCmdSetEvent2-dstStageMask-09392  
  The `dstStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** not include `VK_PIPELINE_STAGE_2_HOST_BIT`

- <a href="#VUID-vkCmdSetEvent2-commandBuffer-03826"
  id="VUID-vkCmdSetEvent2-commandBuffer-03826"></a>
  VUID-vkCmdSetEvent2-commandBuffer-03826  
  The current device mask of `commandBuffer` **must** include exactly
  one physical device

- <a href="#VUID-vkCmdSetEvent2-srcStageMask-03827"
  id="VUID-vkCmdSetEvent2-srcStageMask-03827"></a>
  VUID-vkCmdSetEvent2-srcStageMask-03827  
  The `srcStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** only include pipeline stages valid for the
  queue family that was used to create the command pool that
  `commandBuffer` was allocated from

- <a href="#VUID-vkCmdSetEvent2-dstStageMask-03828"
  id="VUID-vkCmdSetEvent2-dstStageMask-03828"></a>
  VUID-vkCmdSetEvent2-dstStageMask-03828  
  The `dstStageMask` member of any element of the `pMemoryBarriers`,
  `pBufferMemoryBarriers`, or `pImageMemoryBarriers` members of
  `pDependencyInfo` **must** only include pipeline stages valid for the
  queue family that was used to create the command pool that
  `commandBuffer` was allocated from

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetEvent2-commandBuffer-parameter"
  id="VUID-vkCmdSetEvent2-commandBuffer-parameter"></a>
  VUID-vkCmdSetEvent2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetEvent2-event-parameter"
  id="VUID-vkCmdSetEvent2-event-parameter"></a>
  VUID-vkCmdSetEvent2-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkCmdSetEvent2-pDependencyInfo-parameter"
  id="VUID-vkCmdSetEvent2-pDependencyInfo-parameter"></a>
  VUID-vkCmdSetEvent2-pDependencyInfo-parameter  
  `pDependencyInfo` **must** be a valid pointer to a valid
  [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html) structure

- <a href="#VUID-vkCmdSetEvent2-commandBuffer-recording"
  id="VUID-vkCmdSetEvent2-commandBuffer-recording"></a>
  VUID-vkCmdSetEvent2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetEvent2-commandBuffer-cmdpool"
  id="VUID-vkCmdSetEvent2-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetEvent2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdSetEvent2-renderpass"
  id="VUID-vkCmdSetEvent2-renderpass"></a>
  VUID-vkCmdSetEvent2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdSetEvent2-commonparent"
  id="VUID-vkCmdSetEvent2-commonparent"></a>
  VUID-vkCmdSetEvent2-commonparent  
  Both of `commandBuffer`, and `event` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<tr>
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
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetEvent2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
