# VkPhysicalDevicePortabilitySubsetFeaturesKHR(3) Manual Page

## Name

VkPhysicalDevicePortabilitySubsetFeaturesKHR - Structure describing the features that may not be supported by an implementation of the Vulkan 1.0 Portability Subset



## [](#_c_specification)C Specification

The `VkPhysicalDevicePortabilitySubsetFeaturesKHR` structure is defined as:

```c++
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

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`constantAlphaColorBlendFactors` indicates whether this implementation supports constant *alpha* [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blendfactors](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blendfactors) used as source or destination *color* [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blending](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blending).
- []()`events` indicates whether this implementation supports synchronization using [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-events](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-events).
- []()`imageViewFormatReinterpretation` indicates whether this implementation supports a `VkImageView` being created with a texel format containing a different number of components, or a different number of bits in each component, than the texel format of the underlying `VkImage`.
- []()`imageViewFormatSwizzle` indicates whether this implementation supports remapping format components using [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`components`.
- []()`imageView2DOn3DImage` indicates whether this implementation supports a `VkImage` being created with the `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` flag set, permitting a 2D or 2D array image view to be created on a 3D `VkImage`.
- []()`multisampleArrayImage` indicates whether this implementation supports a `VkImage` being created as a 2D array with multiple samples per texel.
- []()`mutableComparisonSamplers` indicates whether this implementation allows descriptors with comparison samplers to be [updated](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-updates).
- []()`pointPolygons` indicates whether this implementation supports [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast) using a *point* [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygonmode).
- []()`samplerMipLodBias` indicates whether this implementation supports setting a [mipmap LOD bias value](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-mipLodBias) when [creating a sampler](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers).
- []()`separateStencilMaskRef` indicates whether this implementation supports separate front and back [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil) reference values.
- []()`shaderSampleRateInterpolationFunctions` indicates whether this implementation supports fragment shaders which use the [`InterpolationFunction`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-InterpolationFunction) capability and the extended instructions `InterpolateAtCentroid`, `InterpolateAtOffset`, and `InterpolateAtSample` from the `GLSL.std.450` extended instruction set. This member is only meaningful if the [`sampleRateShading`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sampleRateShading) feature is supported.
- []()`tessellationIsolines` indicates whether this implementation supports [isoline output](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation-isoline-tessellation) from the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation) stage of a graphics pipeline. This member is only meaningful if the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is supported.
- []()`tessellationPointMode` indicates whether this implementation supports [point output](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation-point-mode) from the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation) stage of a graphics pipeline. This member is only meaningful if the [`tessellationShader`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tessellationShader) feature is supported.
- []()`triangleFans` indicates whether this implementation supports [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-fans](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-fans) primitive topology.
- []()`vertexAttributeAccessBeyondStride` indicates whether this implementation supports accessing a vertex input attribute beyond the stride of the corresponding vertex input binding.

## [](#_description)Description

If the `VkPhysicalDevicePortabilitySubsetFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePortabilitySubsetFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePortabilitySubsetFeaturesKHR-sType-sType)VUID-VkPhysicalDevicePortabilitySubsetFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_portability\_subset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_portability_subset.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePortabilitySubsetFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0