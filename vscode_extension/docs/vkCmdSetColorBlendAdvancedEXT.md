# vkCmdSetColorBlendAdvancedEXT(3) Manual Page

## Name

vkCmdSetColorBlendAdvancedEXT - Specify the advanced color blend state
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the advanced blend
state, call:

``` c
// Provided by VK_EXT_blend_operation_advanced with VK_EXT_extended_dynamic_state3, VK_EXT_blend_operation_advanced with VK_EXT_shader_object
void vkCmdSetColorBlendAdvancedEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorBlendAdvancedEXT*              pColorBlendAdvanced);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstAttachment` the first color attachment the advanced blend
  parameters apply to.

- `attachmentCount` the number of
  [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html) elements in
  the `pColorBlendAdvanced` array.

- `pColorBlendAdvanced` an array of
  [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html) structs that
  specify the advanced color blend parameters for the corresponding
  attachments.

## <a href="#_description" class="anchor"></a>Description

This command sets the advanced blend operation parameters of the
specified attachments for subsequent drawing commands when drawing using
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`srcPremultiplied`,
[VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`dstPremultiplied`,
and
[VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetColorBlendAdvancedEXT-None-09423"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-None-09423"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3ColorBlendAdvanced`](#features-extendedDynamicState3ColorBlendAdvanced)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetColorBlendAdvancedEXT-pColorBlendAdvanced-parameter"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-pColorBlendAdvanced-parameter"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-pColorBlendAdvanced-parameter  
  `pColorBlendAdvanced` **must** be a valid pointer to an array of
  `attachmentCount` valid
  [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html) structures

- <a href="#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-recording"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetColorBlendAdvancedEXT-videocoding"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-videocoding"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetColorBlendAdvancedEXT-attachmentCount-arraylength"
  id="VUID-vkCmdSetColorBlendAdvancedEXT-attachmentCount-arraylength"></a>
  VUID-vkCmdSetColorBlendAdvancedEXT-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`

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

[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html),
[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetColorBlendAdvancedEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
