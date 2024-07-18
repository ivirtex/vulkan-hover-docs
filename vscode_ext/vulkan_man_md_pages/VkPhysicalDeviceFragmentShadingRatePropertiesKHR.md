# VkPhysicalDeviceFragmentShadingRatePropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRatePropertiesKHR - Structure describing
variable fragment shading rate limits that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentShadingRatePropertiesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkPhysicalDeviceFragmentShadingRatePropertiesKHR {
    VkStructureType          sType;
    void*                    pNext;
    VkExtent2D               minFragmentShadingRateAttachmentTexelSize;
    VkExtent2D               maxFragmentShadingRateAttachmentTexelSize;
    uint32_t                 maxFragmentShadingRateAttachmentTexelSizeAspectRatio;
    VkBool32                 primitiveFragmentShadingRateWithMultipleViewports;
    VkBool32                 layeredShadingRateAttachments;
    VkBool32                 fragmentShadingRateNonTrivialCombinerOps;
    VkExtent2D               maxFragmentSize;
    uint32_t                 maxFragmentSizeAspectRatio;
    uint32_t                 maxFragmentShadingRateCoverageSamples;
    VkSampleCountFlagBits    maxFragmentShadingRateRasterizationSamples;
    VkBool32                 fragmentShadingRateWithShaderDepthStencilWrites;
    VkBool32                 fragmentShadingRateWithSampleMask;
    VkBool32                 fragmentShadingRateWithShaderSampleMask;
    VkBool32                 fragmentShadingRateWithConservativeRasterization;
    VkBool32                 fragmentShadingRateWithFragmentShaderInterlock;
    VkBool32                 fragmentShadingRateWithCustomSampleLocations;
    VkBool32                 fragmentShadingRateStrictMultiplyCombiner;
} VkPhysicalDeviceFragmentShadingRatePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-minFragmentShadingRateAttachmentTexelSize"></span>
  `minFragmentShadingRateAttachmentTexelSize` indicates minimum
  supported width and height of the portion of the framebuffer
  corresponding to each texel in a fragment shading rate attachment.
  Each value **must** be less than or equal to the values in
  `maxFragmentShadingRateAttachmentTexelSize`. Each value **must** be a
  power-of-two. It **must** be (0,0) if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not supported.

- <span id="limits-maxFragmentShadingRateAttachmentTexelSize"></span>
  `maxFragmentShadingRateAttachmentTexelSize` indicates maximum
  supported width and height of the portion of the framebuffer
  corresponding to each texel in a fragment shading rate attachment.
  Each value **must** be greater than or equal to the values in
  `minFragmentShadingRateAttachmentTexelSize`. Each value **must** be a
  power-of-two. It **must** be (0,0) if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not supported.

- <span id="limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio"></span>
  `maxFragmentShadingRateAttachmentTexelSizeAspectRatio` indicates the
  maximum ratio between the width and height of the portion of the
  framebuffer corresponding to each texel in a fragment shading rate
  attachment. `maxFragmentShadingRateAttachmentTexelSizeAspectRatio`
  **must** be a power-of-two value, and **must** be less than or equal
  to max(`maxFragmentShadingRateAttachmentTexelSize.width` /
  `minFragmentShadingRateAttachmentTexelSize.height`,
  `maxFragmentShadingRateAttachmentTexelSize.height` /
  `minFragmentShadingRateAttachmentTexelSize.width`). It **must** be 0
  if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not supported.

- <span id="limits-primitiveFragmentShadingRateWithMultipleViewports"></span>
  `primitiveFragmentShadingRateWithMultipleViewports` specifies whether
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
  target="_blank" rel="noopener">primitive fragment shading rate</a>
  **can** be used when multiple viewports are used. If this value is
  `VK_FALSE`, only a single viewport **must** be used, and applications
  **must** not write to the `ViewportMaskNV` or `ViewportIndex` built-in
  when setting `PrimitiveShadingRateKHR`. It **must** be `VK_FALSE` if
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputViewportIndex"
  target="_blank"
  rel="noopener"><code>shaderOutputViewportIndex</code></a> feature, the
  [`VK_EXT_shader_viewport_index_layer`](VK_EXT_shader_viewport_index_layer.html)
  extension, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not supported, or if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  is not supported.

- <span id="limits-layeredShadingRateAttachments"></span>
  `layeredShadingRateAttachments` specifies whether a shading rate
  attachment image view **can** be created with multiple layers. If this
  value is `VK_FALSE`, when creating an image view with a `usage` that
  includes `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`,
  `layerCount` **must** be `1`. It **must** be `VK_FALSE` if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature, the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputViewportIndex"
  target="_blank"
  rel="noopener"><code>shaderOutputViewportIndex</code></a> feature, the
  [`VK_EXT_shader_viewport_index_layer`](VK_EXT_shader_viewport_index_layer.html)
  extension, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
  target="_blank" rel="noopener"><code>geometryShader</code></a> feature
  is not supported, or if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not supported.

- <span id="limits-fragmentShadingRateNonTrivialCombinerOps"></span>
  `fragmentShadingRateNonTrivialCombinerOps` specifies whether
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  enums other than `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` or
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR` **can** be used. It
  **must** be `VK_FALSE` unless either the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is supported.

- <span id="limits-maxFragmentSize"></span> `maxFragmentSize` indicates
  the maximum supported width and height of a fragment. Its `width` and
  `height` members **must** both be power-of-two values. This limit is
  purely informational, and is not validated.

- <span id="limits-maxFragmentSizeAspectRatio"></span>
  `maxFragmentSizeAspectRatio` indicates the maximum ratio between the
  width and height of a fragment. `maxFragmentSizeAspectRatio` **must**
  be a power-of-two value, and **must** be less than or equal to the
  maximum of the `width` and `height` members of `maxFragmentSize`. This
  limit is purely informational, and is not validated.

- <span id="limits-maxFragmentShadingRateCoverageSamples"></span>
  `maxFragmentShadingRateCoverageSamples` specifies the maximum number
  of coverage samples supported in a single fragment.
  `maxFragmentShadingRateCoverageSamples` **must** be less than or equal
  to the product of the `width` and `height` members of
  `maxFragmentSize`, and the sample count reported by
  `maxFragmentShadingRateRasterizationSamples`.
  `maxFragmentShadingRateCoverageSamples` **must** be less than or equal
  to `maxSampleMaskWords` × 32 if
  `fragmentShadingRateWithShaderSampleMask` is supported. This limit is
  purely informational, and is not validated.

- <span id="limits-maxFragmentShadingRateRasterizationSamples"></span>
  `maxFragmentShadingRateRasterizationSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value specifying
  the maximum sample rate supported when a fragment covers multiple
  pixels. This limit is purely informational, and is not validated.

- <span id="limits-fragmentShadingRateWithShaderDepthStencilWrites"></span>
  `fragmentShadingRateWithShaderDepthStencilWrites` specifies whether
  the implementation supports writing `FragDepth` or `FragStencilRefEXT`
  from a fragment shader for multi-pixel fragments. If this value is
  `VK_FALSE`, writing to those built-ins will clamp the fragment shading
  rate to (1,1).

- <span id="limits-fragmentShadingRateWithSampleMask"></span>
  `fragmentShadingRateWithSampleMask` specifies whether the
  implementation supports setting valid bits of
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`pSampleMask`
  to `0` for multi-pixel fragments. If this value is `VK_FALSE`, zeroing
  valid bits in the sample mask will clamp the fragment shading rate to
  (1,1).

- <span id="limits-fragmentShadingRateWithShaderSampleMask"></span>
  `fragmentShadingRateWithShaderSampleMask` specifies whether the
  implementation supports reading or writing `SampleMask` for
  multi-pixel fragments. If this value is `VK_FALSE`, using that
  built-in will clamp the fragment shading rate to (1,1).

- <span id="limits-fragmentShadingRateWithConservativeRasterization"></span>
  `fragmentShadingRateWithConservativeRasterization` specifies whether
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-conservativeraster"
  target="_blank" rel="noopener">conservative rasterization</a> is
  supported for multi-pixel fragments. It **must** be `VK_FALSE` if
  [`VK_EXT_conservative_rasterization`](VK_EXT_conservative_rasterization.html)
  is not supported. If this value is `VK_FALSE`, using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-conservativeraster"
  target="_blank" rel="noopener">conservative rasterization</a> will
  clamp the fragment shading rate to (1,1).

- <span id="limits-fragmentShadingRateWithFragmentShaderInterlock"></span>
  `fragmentShadingRateWithFragmentShaderInterlock` specifies whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-interlock"
  target="_blank" rel="noopener">fragment shader interlock</a> is
  supported for multi-pixel fragments. It **must** be `VK_FALSE` if
  [`VK_EXT_fragment_shader_interlock`](VK_EXT_fragment_shader_interlock.html)
  is not supported. If this value is `VK_FALSE`, using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-interlock"
  target="_blank" rel="noopener">fragment shader interlock</a> will
  clamp the fragment shading rate to (1,1).

- <span id="limits-fragmentShadingRateWithCustomSampleLocations"></span>
  `fragmentShadingRateWithCustomSampleLocations` specifies whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-samplelocations"
  target="_blank" rel="noopener">custom sample locations</a> are
  supported for multi-pixel fragments. It **must** be `VK_FALSE` if
  [`VK_EXT_sample_locations`](VK_EXT_sample_locations.html) is not
  supported. If this value is `VK_FALSE`, using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-samplelocations"
  target="_blank" rel="noopener">custom sample locations</a> will clamp
  the fragment shading rate to (1,1).

- <span id="limits-fragmentShadingRateStrictMultiplyCombiner"></span>
  `fragmentShadingRateStrictMultiplyCombiner` specifies whether
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_MUL_KHR` accurately performs a
  multiplication or not. Implementations where this value is `VK_FALSE`
  will instead combine rates with an addition. If
  `fragmentShadingRateNonTrivialCombinerOps` is `VK_FALSE`,
  implementations **must** report this as `VK_FALSE`. If
  `fragmentShadingRateNonTrivialCombinerOps` is `VK_TRUE`,
  implementations **should** report this as `VK_TRUE`.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Multiplication of the combiner rates using the fragment width/height
in linear space is equivalent to an addition of those values in log2
space. Some implementations inadvertently implemented an addition in
linear space due to unclear requirements originating outside of this
specification. This resulted in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateStrictMultiplyCombiner"
target="_blank"
rel="noopener"><code>fragmentShadingRateStrictMultiplyCombiner</code></a>
being added. Fortunately, this only affects situations where a rate of 1
in either dimension is combined with another rate of 1. All other
combinations result in the exact same result as if multiplication was
performed in linear space due to the clamping logic, and the fact that
both the sum and product of 2 and 2 are equal. In many cases, this limit
will not affect the correct operation of applications.</p></td>
</tr>
</tbody>
</table>

If the `VkPhysicalDeviceFragmentShadingRatePropertiesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These properties are related to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate"
target="_blank" rel="noopener">fragment shading rates</a>.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentShadingRatePropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentShadingRatePropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentShadingRatePropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentShadingRatePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
