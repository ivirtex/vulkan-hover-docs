# VkPhysicalDevicePortabilitySubsetFeaturesKHR(3) Manual Page

## Name

VkPhysicalDevicePortabilitySubsetFeaturesKHR - Structure describing the
features that may not be supported by an implementation of the Vulkan
1.0 Portability Subset



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePortabilitySubsetFeaturesKHR` structure is defined
as:

``` c
// Provided by VK_KHR_portability_subset
typedef struct VkPhysicalDevicePortabilitySubsetFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           constantAlphaColorBlendFactors;
    VkBool32           events;
    VkBool32           imageViewFormatReinterpretation;
    VkBool32           imageViewFormatSwizzle;
    VkBool32           imageView2DOn3DImage;
    VkBool32           multisampleArrayImage;
    VkBool32           mutableComparisonSamplers;
    VkBool32           pointPolygons;
    VkBool32           samplerMipLodBias;
    VkBool32           separateStencilMaskRef;
    VkBool32           shaderSampleRateInterpolationFunctions;
    VkBool32           tessellationIsolines;
    VkBool32           tessellationPointMode;
    VkBool32           triangleFans;
    VkBool32           vertexAttributeAccessBeyondStride;
} VkPhysicalDevicePortabilitySubsetFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-constantAlphaColorBlendFactors"></span>
  `constantAlphaColorBlendFactors` indicates whether this implementation
  supports constant *alpha* <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blendfactors"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blendfactors</a>
  used as source or destination *color* <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blending"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blending</a>.

- <span id="features-events"></span> `events` indicates whether this
  implementation supports synchronization using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-events"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-events</a>.

- <span id="features-imageViewFormatReinterpretation"></span>
  `imageViewFormatReinterpretation` indicates whether this
  implementation supports a `VkImageView` being created with a texel
  format containing a different number of components, or a different
  number of bits in each component, than the texel format of the
  underlying `VkImage`.

- <span id="features-imageViewFormatSwizzle"></span>
  `imageViewFormatSwizzle` indicates whether this implementation
  supports remapping format components using
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`components`.

- <span id="features-imageView2DOn3DImage"></span>
  `imageView2DOn3DImage` indicates whether this implementation supports
  a `VkImage` being created with the
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` flag set, permitting a 2D or
  2D array image view to be created on a 3D `VkImage`.

- <span id="features-multisampleArrayImage"></span>
  `multisampleArrayImage` indicates whether this implementation supports
  a `VkImage` being created as a 2D array with multiple samples per
  texel.

- <span id="features-mutableComparisonSamplers"></span>
  `mutableComparisonSamplers` indicates whether this implementation
  allows descriptors with comparison samplers to be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-updates"
  target="_blank" rel="noopener">updated</a>.

- <span id="features-pointPolygons"></span> `pointPolygons` indicates
  whether this implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast</a>
  using a *point* <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode</a>.

- <span id="features-samplerMipLodBias"></span> `samplerMipLodBias`
  indicates whether this implementation supports setting a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-mipLodBias"
  target="_blank" rel="noopener">mipmap LOD bias value</a> when <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers"
  target="_blank" rel="noopener">creating a sampler</a>.

- <span id="features-separateStencilMaskRef"></span>
  `separateStencilMaskRef` indicates whether this implementation
  supports separate front and back <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil</a>
  reference values.

- <span id="features-shaderSampleRateInterpolationFunctions"></span>
  `shaderSampleRateInterpolationFunctions` indicates whether this
  implementation supports fragment shaders which use the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-InterpolationFunction"
  target="_blank" rel="noopener"><code>InterpolationFunction</code></a>
  capability and the extended instructions `InterpolateAtCentroid`,
  `InterpolateAtOffset`, and `InterpolateAtSample` from the
  `GLSL.std.450` extended instruction set. This member is only
  meaningful if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sampleRateShading"
  target="_blank" rel="noopener"><code>sampleRateShading</code></a>
  feature is supported.

- <span id="features-tessellationIsolines"></span>
  `tessellationIsolines` indicates whether this implementation supports
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation-isoline-tessellation"
  target="_blank" rel="noopener">isoline output</a> from the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation</a>
  stage of a graphics pipeline. This member is only meaningful if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a> are
  supported.

- <span id="features-tessellationPointMode"></span>
  `tessellationPointMode` indicates whether this implementation supports
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation-point-mode"
  target="_blank" rel="noopener">point output</a> from the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation</a>
  stage of a graphics pipeline. This member is only meaningful if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a> are
  supported.

- <span id="features-triangleFans"></span> `triangleFans` indicates
  whether this implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-fans"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-triangle-fans</a>
  primitive topology.

- <span id="features-vertexAttributeAccessBeyondStride"></span>
  `vertexAttributeAccessBeyondStride` indicates whether this
  implementation supports accessing a vertex input attribute beyond the
  stride of the corresponding vertex input binding.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePortabilitySubsetFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePortabilitySubsetFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDevicePortabilitySubsetFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDevicePortabilitySubsetFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDevicePortabilitySubsetFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_portability_subset](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_portability_subset.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePortabilitySubsetFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
