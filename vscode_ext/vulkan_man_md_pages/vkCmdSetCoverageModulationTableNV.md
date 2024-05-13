# vkCmdSetCoverageModulationTableNV(3) Manual Page

## Name

vkCmdSetCoverageModulationTableNV - Specify the coverage modulation
table dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`pCoverageModulationTable` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples, VK_EXT_shader_object with VK_NV_framebuffer_mixed_samples
void vkCmdSetCoverageModulationTableNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    coverageModulationTableCount,
    const float*                                pCoverageModulationTable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `coverageModulationTableCount` specifies the number of elements in
  `pCoverageModulationTable`.

- `pCoverageModulationTable` specifies the table of modulation factors
  containing a value for each number of covered samples.

## <a href="#_description" class="anchor"></a>Description

This command sets the table of modulation factors for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationTableCount`,
and
[VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`pCoverageModulationTable`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetCoverageModulationTableNV-None-09423"
  id="VUID-vkCmdSetCoverageModulationTableNV-None-09423"></a>
  VUID-vkCmdSetCoverageModulationTableNV-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3CoverageModulationTable`](#features-extendedDynamicState3CoverageModulationTable)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-parameter"
  id="VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetCoverageModulationTableNV-pCoverageModulationTable-parameter"
  id="VUID-vkCmdSetCoverageModulationTableNV-pCoverageModulationTable-parameter"></a>
  VUID-vkCmdSetCoverageModulationTableNV-pCoverageModulationTable-parameter  
  `pCoverageModulationTable` **must** be a valid pointer to an array of
  `coverageModulationTableCount` `float` values

- <a
  href="#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-recording"
  id="VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-recording"></a>
  VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetCoverageModulationTableNV-videocoding"
  id="VUID-vkCmdSetCoverageModulationTableNV-videocoding"></a>
  VUID-vkCmdSetCoverageModulationTableNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetCoverageModulationTableNV-coverageModulationTableCount-arraylength"
  id="VUID-vkCmdSetCoverageModulationTableNV-coverageModulationTableCount-arraylength"></a>
  VUID-vkCmdSetCoverageModulationTableNV-coverageModulationTableCount-arraylength  
  `coverageModulationTableCount` **must** be greater than `0`

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
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetCoverageModulationTableNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
