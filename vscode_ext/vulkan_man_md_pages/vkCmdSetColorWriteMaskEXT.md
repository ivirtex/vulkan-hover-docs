# vkCmdSetColorWriteMaskEXT(3) Manual Page

## Name

vkCmdSetColorWriteMaskEXT - Specify the color write masks for each
attachment dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the color write
masks, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetColorWriteMaskEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorComponentFlags*                pColorWriteMasks);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstAttachment` the first color attachment the color write masks
  apply to.

- `attachmentCount` the number of
  [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html) values in the
  `pColorWriteMasks` array.

- `pColorWriteMasks` an array of
  [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html) values that
  specify the color write masks of the corresponding attachments.

## <a href="#_description" class="anchor"></a>Description

This command sets the color write masks of the specified attachments for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`colorWriteMask`
values used to create the currently active pipeline.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Formats with bits that are shared between components specified by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html">VkColorComponentFlagBits</a>, such
as <code>VK_FORMAT_E5B9G9R9_UFLOAT_PACK32</code>, cannot have their
channels individually masked by this functionality; either all
components that share bits have to be enabled, or none of them.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-None-09423"
  id="VUID-vkCmdSetColorWriteMaskEXT-None-09423"></a>
  VUID-vkCmdSetColorWriteMaskEXT-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3ColorWriteMask`](#features-extendedDynamicState3ColorWriteMask)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-pColorWriteMasks-parameter"
  id="VUID-vkCmdSetColorWriteMaskEXT-pColorWriteMasks-parameter"></a>
  VUID-vkCmdSetColorWriteMaskEXT-pColorWriteMasks-parameter  
  `pColorWriteMasks` **must** be a valid pointer to an array of
  `attachmentCount` valid combinations of
  [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html) values

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-recording"
  id="VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-videocoding"
  id="VUID-vkCmdSetColorWriteMaskEXT-videocoding"></a>
  VUID-vkCmdSetColorWriteMaskEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetColorWriteMaskEXT-attachmentCount-arraylength"
  id="VUID-vkCmdSetColorWriteMaskEXT-attachmentCount-arraylength"></a>
  VUID-vkCmdSetColorWriteMaskEXT-attachmentCount-arraylength  
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
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetColorWriteMaskEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
