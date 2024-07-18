# vkCmdSetDiscardRectangleEXT(3) Manual Page

## Name

vkCmdSetDiscardRectangleEXT - Set discard rectangles dynamically for a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the discard
rectangles, call:

``` c
// Provided by VK_EXT_discard_rectangles
void vkCmdSetDiscardRectangleEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstDiscardRectangle,
    uint32_t                                    discardRectangleCount,
    const VkRect2D*                             pDiscardRectangles);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstDiscardRectangle` is the index of the first discard rectangle
  whose state is updated by the command.

- `discardRectangleCount` is the number of discard rectangles whose
  state are updated by the command.

- `pDiscardRectangles` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures specifying discard rectangles.

## <a href="#_description" class="anchor"></a>Description

The discard rectangle taken from element i of `pDiscardRectangles`
replace the current state for the discard rectangle at index
`firstDiscardRectangle` + i, for i in \[0, `discardRectangleCount`).

This command sets the discard rectangles for subsequent drawing commands
when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`pDiscardRectangles`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-firstDiscardRectangle-00585"
  id="VUID-vkCmdSetDiscardRectangleEXT-firstDiscardRectangle-00585"></a>
  VUID-vkCmdSetDiscardRectangleEXT-firstDiscardRectangle-00585  
  The sum of `firstDiscardRectangle` and `discardRectangleCount`
  **must** be less than or equal to
  [VkPhysicalDeviceDiscardRectanglePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html)::`maxDiscardRectangles`

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-x-00587"
  id="VUID-vkCmdSetDiscardRectangleEXT-x-00587"></a>
  VUID-vkCmdSetDiscardRectangleEXT-x-00587  
  The `x` and `y` member of `offset` in each [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)
  element of `pDiscardRectangles` **must** be greater than or equal to
  `0`

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-offset-00588"
  id="VUID-vkCmdSetDiscardRectangleEXT-offset-00588"></a>
  VUID-vkCmdSetDiscardRectangleEXT-offset-00588  
  Evaluation of (`offset.x` + `extent.width`) in each
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) element of `pDiscardRectangles` **must** not
  cause a signed integer addition overflow

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-offset-00589"
  id="VUID-vkCmdSetDiscardRectangleEXT-offset-00589"></a>
  VUID-vkCmdSetDiscardRectangleEXT-offset-00589  
  Evaluation of (`offset.y` + `extent.height`) in each
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) element of `pDiscardRectangles` **must** not
  cause a signed integer addition overflow

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-viewportScissor2D-04788"
  id="VUID-vkCmdSetDiscardRectangleEXT-viewportScissor2D-04788"></a>
  VUID-vkCmdSetDiscardRectangleEXT-viewportScissor2D-04788  
  If this command is recorded in a secondary command buffer with
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D`
  enabled, then this function **must** not be called

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-pDiscardRectangles-parameter"
  id="VUID-vkCmdSetDiscardRectangleEXT-pDiscardRectangles-parameter"></a>
  VUID-vkCmdSetDiscardRectangleEXT-pDiscardRectangles-parameter  
  `pDiscardRectangles` **must** be a valid pointer to an array of
  `discardRectangleCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-recording"
  id="VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetDiscardRectangleEXT-videocoding"
  id="VUID-vkCmdSetDiscardRectangleEXT-videocoding"></a>
  VUID-vkCmdSetDiscardRectangleEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetDiscardRectangleEXT-discardRectangleCount-arraylength"
  id="VUID-vkCmdSetDiscardRectangleEXT-discardRectangleCount-arraylength"></a>
  VUID-vkCmdSetDiscardRectangleEXT-discardRectangleCount-arraylength  
  `discardRectangleCount` **must** be greater than `0`

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

[VK_EXT_discard_rectangles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDiscardRectangleEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
