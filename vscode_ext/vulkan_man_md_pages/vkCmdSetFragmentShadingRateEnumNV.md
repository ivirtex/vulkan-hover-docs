# vkCmdSetFragmentShadingRateEnumNV(3) Manual Page

## Name

vkCmdSetFragmentShadingRateEnumNV - Set pipeline fragment shading rate
dynamically for a command buffer using enums



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the pipeline fragment
shading rate and combiner operation, call:

``` c
// Provided by VK_NV_fragment_shading_rate_enums
void vkCmdSetFragmentShadingRateEnumNV(
    VkCommandBuffer                             commandBuffer,
    VkFragmentShadingRateNV                     shadingRate,
    const VkFragmentShadingRateCombinerOpKHR    combinerOps[2]);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `shadingRate` specifies a
  [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html) enum
  indicating the pipeline fragment shading rate for subsequent drawing
  commands.

- `combinerOps` specifies a
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  determining how the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-pipeline"
  target="_blank" rel="noopener">pipeline</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
  target="_blank" rel="noopener">primitive</a>, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">attachment shading rates</a> are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-combining"
  target="_blank" rel="noopener">combined</a> for fragments generated by
  subsequent drawing commands.

## <a href="#_description" class="anchor"></a>Description

This command sets the pipeline fragment shading rate and combiner
operation for subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`
set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
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
<p>This command allows specifying additional shading rates beyond those
supported by <a
href="vkCmdSetFragmentShadingRateKHR.html">vkCmdSetFragmentShadingRateKHR</a>.
For more information, refer to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html"><code>VK_NV_fragment_shading_rate_enums</code></a>
appendix.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04576"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04576"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04576  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a> is not
  enabled, `shadingRate` **must** be
  `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV`

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-supersampleFragmentShadingRates-04577"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-supersampleFragmentShadingRates-04577"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-supersampleFragmentShadingRates-04577  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-supersampleFragmentShadingRates"
  target="_blank"
  rel="noopener"><code>supersampleFragmentShadingRates</code></a> is not
  enabled, `shadingRate` **must** not be
  `VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV`,
  `VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV`,
  `VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV`, or
  `VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV`

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-noInvocationFragmentShadingRates-04578"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-noInvocationFragmentShadingRates-04578"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-noInvocationFragmentShadingRates-04578  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-noInvocationFragmentShadingRates"
  target="_blank"
  rel="noopener"><code>noInvocationFragmentShadingRates</code></a> is
  not enabled, `shadingRate` **must** not be
  `VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV`

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentShadingRateEnums-04579"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentShadingRateEnums-04579"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentShadingRateEnums-04579  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentShadingRateEnums"
  target="_blank" rel="noopener"><code>fragmentShadingRateEnums</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04580"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04580"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-pipelineFragmentShadingRate-04580  
  One of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> features
  **must** be enabled

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-primitiveFragmentShadingRate-04581"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-primitiveFragmentShadingRate-04581"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-primitiveFragmentShadingRate-04581  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  is not enabled, `combinerOps`\[0\] **must** be
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-attachmentFragmentShadingRate-04582"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-attachmentFragmentShadingRate-04582"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-attachmentFragmentShadingRate-04582  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not enabled, `combinerOps`\[1\] **must** be
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentSizeNonTrivialCombinerOps-04583"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentSizeNonTrivialCombinerOps-04583"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-fragmentSizeNonTrivialCombinerOps-04583  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateNonTrivialCombinerOps"
  target="_blank"
  rel="noopener"><code>fragmentSizeNonTrivialCombinerOps</code></a>
  limit is not supported, elements of `combinerOps` **must** be either
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` or
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-parameter"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetFragmentShadingRateEnumNV-shadingRate-parameter"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-shadingRate-parameter"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-shadingRate-parameter  
  `shadingRate` **must** be a valid
  [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html) value

- <a href="#VUID-vkCmdSetFragmentShadingRateEnumNV-combinerOps-parameter"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-combinerOps-parameter"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-combinerOps-parameter  
  Each element of `combinerOps` **must** be a valid
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  value

- <a
  href="#VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-recording"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-recording"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetFragmentShadingRateEnumNV-videocoding"
  id="VUID-vkCmdSetFragmentShadingRateEnumNV-videocoding"></a>
  VUID-vkCmdSetFragmentShadingRateEnumNV-videocoding  
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

[VK_NV_fragment_shading_rate_enums](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html),
[VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetFragmentShadingRateEnumNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700