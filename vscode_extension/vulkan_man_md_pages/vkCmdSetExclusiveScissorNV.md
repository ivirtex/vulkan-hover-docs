# vkCmdSetExclusiveScissorNV(3) Manual Page

## Name

vkCmdSetExclusiveScissorNV - Set exclusive scissor rectangles
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the exclusive scissor
rectangles, call:

``` c
// Provided by VK_NV_scissor_exclusive
void vkCmdSetExclusiveScissorNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstExclusiveScissor,
    uint32_t                                    exclusiveScissorCount,
    const VkRect2D*                             pExclusiveScissors);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstExclusiveScissor` is the index of the first exclusive scissor
  rectangle whose state is updated by the command.

- `exclusiveScissorCount` is the number of exclusive scissor rectangles
  updated by the command.

- `pExclusiveScissors` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures defining exclusive scissor
  rectangles.

## <a href="#_description" class="anchor"></a>Description

The scissor rectangles taken from element i of `pExclusiveScissors`
replace the current state for the scissor index
`firstExclusiveScissor` + i, for i in \[0, `exclusiveScissorCount`).

This command sets the exclusive scissor rectangles for subsequent
drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)::`pExclusiveScissors`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetExclusiveScissorNV-None-02031"
  id="VUID-vkCmdSetExclusiveScissorNV-None-02031"></a>
  VUID-vkCmdSetExclusiveScissorNV-None-02031  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-exclusiveScissor"
  target="_blank" rel="noopener"><code>exclusiveScissor</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02034"
  id="VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02034"></a>
  VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02034  
  The sum of `firstExclusiveScissor` and `exclusiveScissorCount`
  **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`,
  inclusive

- <a href="#VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02035"
  id="VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02035"></a>
  VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02035  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `firstExclusiveScissor` **must** be `0`

- <a href="#VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-02036"
  id="VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-02036"></a>
  VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-02036  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `exclusiveScissorCount` **must** be `1`

- <a href="#VUID-vkCmdSetExclusiveScissorNV-x-02037"
  id="VUID-vkCmdSetExclusiveScissorNV-x-02037"></a>
  VUID-vkCmdSetExclusiveScissorNV-x-02037  
  The `x` and `y` members of `offset` in each member of
  `pExclusiveScissors` **must** be greater than or equal to `0`

- <a href="#VUID-vkCmdSetExclusiveScissorNV-offset-02038"
  id="VUID-vkCmdSetExclusiveScissorNV-offset-02038"></a>
  VUID-vkCmdSetExclusiveScissorNV-offset-02038  
  Evaluation of (`offset.x` + `extent.width`) for each member of
  `pExclusiveScissors` **must** not cause a signed integer addition
  overflow

- <a href="#VUID-vkCmdSetExclusiveScissorNV-offset-02039"
  id="VUID-vkCmdSetExclusiveScissorNV-offset-02039"></a>
  VUID-vkCmdSetExclusiveScissorNV-offset-02039  
  Evaluation of (`offset.y` + `extent.height`) for each member of
  `pExclusiveScissors` **must** not cause a signed integer addition
  overflow

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-parameter"
  id="VUID-vkCmdSetExclusiveScissorNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetExclusiveScissorNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetExclusiveScissorNV-pExclusiveScissors-parameter"
  id="VUID-vkCmdSetExclusiveScissorNV-pExclusiveScissors-parameter"></a>
  VUID-vkCmdSetExclusiveScissorNV-pExclusiveScissors-parameter  
  `pExclusiveScissors` **must** be a valid pointer to an array of
  `exclusiveScissorCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-recording"
  id="VUID-vkCmdSetExclusiveScissorNV-commandBuffer-recording"></a>
  VUID-vkCmdSetExclusiveScissorNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetExclusiveScissorNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetExclusiveScissorNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetExclusiveScissorNV-videocoding"
  id="VUID-vkCmdSetExclusiveScissorNV-videocoding"></a>
  VUID-vkCmdSetExclusiveScissorNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-arraylength"
  id="VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-arraylength"></a>
  VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-arraylength  
  `exclusiveScissorCount` **must** be greater than `0`

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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_scissor_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_scissor_exclusive.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetExclusiveScissorNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
