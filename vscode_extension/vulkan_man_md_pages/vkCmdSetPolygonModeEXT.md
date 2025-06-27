# vkCmdSetPolygonModeEXT(3) Manual Page

## Name

vkCmdSetPolygonModeEXT - Specify polygon mode dynamically for a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the polygon mode,
call:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetPolygonModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkPolygonMode                               polygonMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `polygonMode` specifies polygon mode.

## <a href="#_description" class="anchor"></a>Description

This command sets the polygon mode for subsequent drawing commands when
drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_POLYGON_MODE_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`polygonMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetPolygonModeEXT-None-09423"
  id="VUID-vkCmdSetPolygonModeEXT-None-09423"></a>
  VUID-vkCmdSetPolygonModeEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3PolygonMode`](#features-extendedDynamicState3PolygonMode)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

- <a href="#VUID-vkCmdSetPolygonModeEXT-fillModeNonSolid-07424"
  id="VUID-vkCmdSetPolygonModeEXT-fillModeNonSolid-07424"></a>
  VUID-vkCmdSetPolygonModeEXT-fillModeNonSolid-07424  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fillModeNonSolid"
  target="_blank" rel="noopener"><code>fillModeNonSolid</code></a>
  feature is not enabled, `polygonMode` **must** be
  `VK_POLYGON_MODE_FILL` or `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

- <a href="#VUID-vkCmdSetPolygonModeEXT-polygonMode-07425"
  id="VUID-vkCmdSetPolygonModeEXT-polygonMode-07425"></a>
  VUID-vkCmdSetPolygonModeEXT-polygonMode-07425  
  If the [`VK_NV_fill_rectangle`](VK_NV_fill_rectangle.html) extension
  is not enabled, `polygonMode` **must** not be
  `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetPolygonModeEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetPolygonModeEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetPolygonModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetPolygonModeEXT-polygonMode-parameter"
  id="VUID-vkCmdSetPolygonModeEXT-polygonMode-parameter"></a>
  VUID-vkCmdSetPolygonModeEXT-polygonMode-parameter  
  `polygonMode` **must** be a valid [VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html)
  value

- <a href="#VUID-vkCmdSetPolygonModeEXT-commandBuffer-recording"
  id="VUID-vkCmdSetPolygonModeEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetPolygonModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetPolygonModeEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetPolygonModeEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetPolygonModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetPolygonModeEXT-videocoding"
  id="VUID-vkCmdSetPolygonModeEXT-videocoding"></a>
  VUID-vkCmdSetPolygonModeEXT-videocoding  
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
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetPolygonModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
