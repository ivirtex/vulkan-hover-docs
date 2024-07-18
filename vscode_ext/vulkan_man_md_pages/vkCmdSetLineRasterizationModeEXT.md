# vkCmdSetLineRasterizationModeEXT(3) Manual Page

## Name

vkCmdSetLineRasterizationModeEXT - Specify the line rasterization mode
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`lineRasterizationMode` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_line_rasterization, VK_EXT_line_rasterization with VK_EXT_shader_object
void vkCmdSetLineRasterizationModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkLineRasterizationModeEXT                  lineRasterizationMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `lineRasterizationMode` specifies the `lineRasterizationMode` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `lineRasterizationMode` state for subsequent
drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)::`lineRasterizationMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetLineRasterizationModeEXT-None-09423"
  id="VUID-vkCmdSetLineRasterizationModeEXT-None-09423"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3LineRasterizationMode`](#features-extendedDynamicState3LineRasterizationMode)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

- <a
  href="#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07418"
  id="VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07418"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07418  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rectangularLines"
  target="_blank" rel="noopener"><code>rectangularLines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07419"
  id="VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07419"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07419  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bresenhamLines"
  target="_blank" rel="noopener"><code>bresenhamLines</code></a> feature
  **must** be enabled

- <a
  href="#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07420"
  id="VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07420"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07420  
  If `lineRasterizationMode` is
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-smoothLines"
  target="_blank" rel="noopener"><code>smoothLines</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-parameter"
  id="VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-parameter"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-parameter  
  `lineRasterizationMode` **must** be a valid
  [VkLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeEXT.html) value

- <a href="#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-recording"
  id="VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetLineRasterizationModeEXT-videocoding"
  id="VUID-vkCmdSetLineRasterizationModeEXT-videocoding"></a>
  VUID-vkCmdSetLineRasterizationModeEXT-videocoding  
  This command **must** only be called outside of a video coding scope

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

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetLineRasterizationModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
