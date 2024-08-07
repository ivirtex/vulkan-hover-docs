# vkCmdSetConservativeRasterizationModeEXT(3) Manual Page

## Name

vkCmdSetConservativeRasterizationModeEXT - Specify the conservative
rasterization mode dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`conservativeRasterizationMode`, call:

``` c
// Provided by VK_EXT_conservative_rasterization with VK_EXT_extended_dynamic_state3, VK_EXT_conservative_rasterization with VK_EXT_shader_object
void vkCmdSetConservativeRasterizationModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkConservativeRasterizationModeEXT          conservativeRasterizationMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `conservativeRasterizationMode` specifies the
  `conservativeRasterizationMode` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `conservativeRasterizationMode` state for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with
`VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetConservativeRasterizationModeEXT-None-09423"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-None-09423"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3ConservativeRasterizationMode`](#features-extendedDynamicState3ConservativeRasterizationMode)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetConservativeRasterizationModeEXT-conservativeRasterizationMode-parameter"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-conservativeRasterizationMode-parameter"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-conservativeRasterizationMode-parameter  
  `conservativeRasterizationMode` **must** be a valid
  [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConservativeRasterizationModeEXT.html)
  value

- <a
  href="#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-recording"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetConservativeRasterizationModeEXT-videocoding"
  id="VUID-vkCmdSetConservativeRasterizationModeEXT-videocoding"></a>
  VUID-vkCmdSetConservativeRasterizationModeEXT-videocoding  
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

[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html),
[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConservativeRasterizationModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetConservativeRasterizationModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
