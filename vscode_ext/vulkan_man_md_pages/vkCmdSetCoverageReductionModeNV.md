# vkCmdSetCoverageReductionModeNV(3) Manual Page

## Name

vkCmdSetCoverageReductionModeNV - Specify the coverage reduction mode
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the
`coverageReductionMode` state, call:

``` c
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_coverage_reduction_mode, VK_EXT_shader_object with VK_NV_coverage_reduction_mode
void vkCmdSetCoverageReductionModeNV(
    VkCommandBuffer                             commandBuffer,
    VkCoverageReductionModeNV                   coverageReductionMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `coverageReductionMode` specifies the `coverageReductionMode` state.

## <a href="#_description" class="anchor"></a>Description

This command sets the `coverageReductionMode` state for subsequent
drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)::`coverageReductionMode`
value used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetCoverageReductionModeNV-None-09423"
  id="VUID-vkCmdSetCoverageReductionModeNV-None-09423"></a>
  VUID-vkCmdSetCoverageReductionModeNV-None-09423  
  At least one of the following **must** be true:

  - The
    [`extendedDynamicState3CoverageReductionMode`](#features-extendedDynamicState3CoverageReductionMode)
    feature is enabled

  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-parameter"
  id="VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetCoverageReductionModeNV-coverageReductionMode-parameter"
  id="VUID-vkCmdSetCoverageReductionModeNV-coverageReductionMode-parameter"></a>
  VUID-vkCmdSetCoverageReductionModeNV-coverageReductionMode-parameter  
  `coverageReductionMode` **must** be a valid
  [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html) value

- <a href="#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-recording"
  id="VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-recording"></a>
  VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetCoverageReductionModeNV-videocoding"
  id="VUID-vkCmdSetCoverageReductionModeNV-videocoding"></a>
  VUID-vkCmdSetCoverageReductionModeNV-videocoding  
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
[VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetCoverageReductionModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
