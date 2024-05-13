# vkCmdSetColorBlendEquationEXT(3) Manual Page

## Name

vkCmdSetColorBlendEquationEXT - Specify the blend factors and operations
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> color blend factors
and operations, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetColorBlendEquationEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorBlendEquationEXT*              pColorBlendEquations);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstAttachment` the first color attachment the color blend factors
  and operations apply to.

- `attachmentCount` the number of
  [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html) elements in
  the `pColorBlendEquations` array.

- `pColorBlendEquations` an array of
  [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html) structs that
  specify the color blend factors and operations for the corresponding
  attachments.

## <a href="#_description" class="anchor"></a>Description

This command sets the color blending factors and operations of the
specified attachments for subsequent drawing commands when drawing using
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT` set
in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`srcColorBlendFactor`,
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`dstColorBlendFactor`,
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`colorBlendOp`,
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`srcAlphaBlendFactor`,
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`dstAlphaBlendFactor`,
and
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`alphaBlendOp`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetColorBlendEquationEXT-None-09423"
  id="VUID-vkCmdSetColorBlendEquationEXT-None-09423"></a>
  VUID-vkCmdSetColorBlendEquationEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3ColorBlendEquation`](#features-extendedDynamicState3ColorBlendEquation)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetColorBlendEquationEXT-pColorBlendEquations-parameter"
  id="VUID-vkCmdSetColorBlendEquationEXT-pColorBlendEquations-parameter"></a>
  VUID-vkCmdSetColorBlendEquationEXT-pColorBlendEquations-parameter  
  `pColorBlendEquations` **must** be a valid pointer to an array of
  `attachmentCount` valid
  [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html) structures

- <a href="#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-recording"
  id="VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetColorBlendEquationEXT-videocoding"
  id="VUID-vkCmdSetColorBlendEquationEXT-videocoding"></a>
  VUID-vkCmdSetColorBlendEquationEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetColorBlendEquationEXT-attachmentCount-arraylength"
  id="VUID-vkCmdSetColorBlendEquationEXT-attachmentCount-arraylength"></a>
  VUID-vkCmdSetColorBlendEquationEXT-attachmentCount-arraylength  
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

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetColorBlendEquationEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
