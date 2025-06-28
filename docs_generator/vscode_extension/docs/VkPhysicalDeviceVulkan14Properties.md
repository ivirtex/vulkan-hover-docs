# VkPhysicalDeviceVulkan14Properties(3) Manual Page

## Name

VkPhysicalDeviceVulkan14Properties - Structure specifying physical device properties for functionality promoted to Vulkan 1.4



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVulkan14Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceVulkan14Properties {
    VkStructureType                       sType;
    void*                                 pNext;
    uint32_t                              lineSubPixelPrecisionBits;
    uint32_t                              maxVertexAttribDivisor;
    VkBool32                              supportsNonZeroFirstInstance;
    uint32_t                              maxPushDescriptors;
    VkBool32                              dynamicRenderingLocalReadDepthStencilAttachments;
    VkBool32                              dynamicRenderingLocalReadMultisampledAttachments;
    VkBool32                              earlyFragmentMultisampleCoverageAfterSampleCounting;
    VkBool32                              earlyFragmentSampleMaskTestBeforeSampleCounting;
    VkBool32                              depthStencilSwizzleOneSupport;
    VkBool32                              polygonModePointSize;
    VkBool32                              nonStrictSinglePixelWideLinesUseParallelogram;
    VkBool32                              nonStrictWideLinesUseParallelogram;
    VkBool32                              blockTexelViewCompatibleMultipleLayers;
    uint32_t                              maxCombinedImageSamplerDescriptorCount;
    VkBool32                              fragmentShadingRateClampCombinerInputs;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessStorageBuffers;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessUniformBuffers;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessVertexInputs;
    VkPipelineRobustnessImageBehavior     defaultRobustnessImages;
    uint32_t                              copySrcLayoutCount;
    VkImageLayout*                        pCopySrcLayouts;
    uint32_t                              copyDstLayoutCount;
    VkImageLayout*                        pCopyDstLayouts;
    uint8_t                               optimalTilingLayoutUUID[VK_UUID_SIZE];
    VkBool32                              identicalMemoryTypeRequirements;
} VkPhysicalDeviceVulkan14Properties;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`lineSubPixelPrecisionBits` is the number of bits of subpixel precision in framebuffer coordinates xf and yf when rasterizing [line segments](#primsrast-lines).
- []()`maxVertexAttribDivisor` is the maximum value of the number of instances that will repeat the value of vertex attribute data when instanced rendering is enabled.
- []()`supportsNonZeroFirstInstance` specifies whether a non-zero value for the `firstInstance` parameter of [drawing commands](#drawing) is supported when [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html)::`divisor` is not `1`.
- []()`maxPushDescriptors` is the maximum number of descriptors that **can** be used in a descriptor set layout created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT` set.
- []()`dynamicRenderingLocalReadDepthStencilAttachments` is `VK_TRUE` if the implementation supports local reads of depth/stencil attachments, `VK_FALSE` otherwise.
- []()`dynamicRenderingLocalReadMultisampledAttachments` is `VK_TRUE` if the implementation supports local reads of multisampled attachments, `VK_FALSE` otherwise.
- `earlyFragmentMultisampleCoverageAfterSampleCounting` is a boolean value indicating whether the [fragment shading](#fragops-shader) and [multisample coverage](#fragops-covg) operations are performed after [sample counting](#fragops-samplecount) for [fragment shaders](#fragops-shader) with `EarlyFragmentTests` execution mode.
- `earlyFragmentSampleMaskTestBeforeSampleCounting` is a boolean value indicating whether the [sample mask test](#fragops-samplemask) operation is performed before [sample counting](#fragops-samplecount) for [fragment shaders](#fragops-shader) using the `EarlyFragmentTests` execution mode.
- `depthStencilSwizzleOneSupport` is a boolean indicating that depth/stencil texturing operations with `VK_COMPONENT_SWIZZLE_ONE` have defined behavior.
- `polygonModePointSize` is a boolean value indicating whether the point size of the final rasterization of polygons with `VK_POLYGON_MODE_POINT` is controlled by `PointSize`.
- `nonStrictSinglePixelWideLinesUseParallelogram` is a boolean value indicating whether non-strict lines with a width of 1.0 are rasterized as parallelograms or using Bresenham’s algorithm.
- `nonStrictWideLinesUseParallelogram` is a boolean value indicating whether non-strict lines with a width greater than 1.0 are rasterized as parallelograms or using Bresenham’s algorithm.
- `blockTexelViewCompatibleMultipleLayers` is a boolean value indicating that an implementation supports creating image views with `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` where the `layerCount` member of `subresourceRange` is greater than `1`.
- `maxCombinedImageSamplerDescriptorCount` is the maximum number of combined image sampler descriptors that the implementation uses to access any of the [formats that require a sampler Y′CBCR conversion](#formats-requiring-sampler-ycbcr-conversion) supported by the implementation.
- `fragmentShadingRateClampCombinerInputs` is a boolean value indicating that an implementation clamps the inputs to [combiner operations](#primsrast-fragment-shading-rate-combining).
- `defaultRobustnessStorageBuffers` describes the behavior of out of bounds accesses made to storage buffers when no robustness features are enabled
- `defaultRobustnessUniformBuffers` describes the behavior of out of bounds accesses made to uniform buffers when no robustness features are enabled
- `defaultRobustnessVertexInputs` describes the behavior of out of bounds accesses made to vertex input attributes when no robustness features are enabled
- `defaultRobustnessImages` describes the behavior of out of bounds accesses made to images when no robustness features are enabled
- `copySrcLayoutCount` is an integer related to the number of image layouts for host copies from images available or queried, as described below.
- `pCopySrcLayouts` is a pointer to an array of [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) in which supported image layouts for use with host copy operations from images are returned.
- `copyDstLayoutCount` is an integer related to the number of image layouts for host copies to images available or queried, as described below.
- `pCopyDstLayouts` is a pointer to an array of [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) in which supported image layouts for use with host copy operations to images are returned.
- `optimalTilingLayoutUUID` is an array of `VK_UUID_SIZE` `uint8_t` values representing a universally unique identifier for the implementation’s swizzling layout of images created with `VK_IMAGE_TILING_OPTIMAL`.
- `identicalMemoryTypeRequirements` indicates that specifying the `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` flag in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` does not affect the memory type requirements of the image.

If the `VkPhysicalDeviceVulkan14Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties correspond to Vulkan 1.4 functionality.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVulkan14Properties-sType-sType)VUID-VkPhysicalDeviceVulkan14Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_4_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html), [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVulkan14Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0