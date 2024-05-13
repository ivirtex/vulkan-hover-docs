# vkCmdSetScissor(3) Manual Page

## Name

vkCmdSetScissor - Set scissor rectangles dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the scissor
rectangles, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdSetScissor(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstScissor,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstScissor` is the index of the first scissor whose state is
  updated by the command.

- `scissorCount` is the number of scissors whose rectangles are updated
  by the command.

- `pScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)
  structures defining scissor rectangles.

## <a href="#_description" class="anchor"></a>Description

The scissor rectangles taken from element i of `pScissors` replace the
current state for the scissor index `firstScissor` + i, for i in \[0,
`scissorCount`).

This command sets the scissor rectangles for subsequent drawing commands
when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_SCISSOR` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`pScissors`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetScissor-firstScissor-00592"
  id="VUID-vkCmdSetScissor-firstScissor-00592"></a>
  VUID-vkCmdSetScissor-firstScissor-00592  
  The sum of `firstScissor` and `scissorCount` **must** be between `1`
  and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive

- <a href="#VUID-vkCmdSetScissor-firstScissor-00593"
  id="VUID-vkCmdSetScissor-firstScissor-00593"></a>
  VUID-vkCmdSetScissor-firstScissor-00593  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `firstScissor` **must** be `0`

- <a href="#VUID-vkCmdSetScissor-scissorCount-00594"
  id="VUID-vkCmdSetScissor-scissorCount-00594"></a>
  VUID-vkCmdSetScissor-scissorCount-00594  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `scissorCount` **must** be `1`

- <a href="#VUID-vkCmdSetScissor-x-00595"
  id="VUID-vkCmdSetScissor-x-00595"></a> VUID-vkCmdSetScissor-x-00595  
  The `x` and `y` members of `offset` member of any element of
  `pScissors` **must** be greater than or equal to `0`

- <a href="#VUID-vkCmdSetScissor-offset-00596"
  id="VUID-vkCmdSetScissor-offset-00596"></a>
  VUID-vkCmdSetScissor-offset-00596  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-vkCmdSetScissor-offset-00597"
  id="VUID-vkCmdSetScissor-offset-00597"></a>
  VUID-vkCmdSetScissor-offset-00597  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-vkCmdSetScissor-viewportScissor2D-04789"
  id="VUID-vkCmdSetScissor-viewportScissor2D-04789"></a>
  VUID-vkCmdSetScissor-viewportScissor2D-04789  
  If this command is recorded in a secondary command buffer with
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled, then this function **must** not be called

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetScissor-commandBuffer-parameter"
  id="VUID-vkCmdSetScissor-commandBuffer-parameter"></a>
  VUID-vkCmdSetScissor-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetScissor-pScissors-parameter"
  id="VUID-vkCmdSetScissor-pScissors-parameter"></a>
  VUID-vkCmdSetScissor-pScissors-parameter  
  `pScissors` **must** be a valid pointer to an array of `scissorCount`
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-vkCmdSetScissor-commandBuffer-recording"
  id="VUID-vkCmdSetScissor-commandBuffer-recording"></a>
  VUID-vkCmdSetScissor-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetScissor-commandBuffer-cmdpool"
  id="VUID-vkCmdSetScissor-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetScissor-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetScissor-videocoding"
  id="VUID-vkCmdSetScissor-videocoding"></a>
  VUID-vkCmdSetScissor-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetScissor-scissorCount-arraylength"
  id="VUID-vkCmdSetScissor-scissorCount-arraylength"></a>
  VUID-vkCmdSetScissor-scissorCount-arraylength  
  `scissorCount` **must** be greater than `0`

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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetScissor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
