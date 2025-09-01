# VkPhysicalDeviceExtendedDynamicState3FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicState3FeaturesEXT - Structure describing what extended dynamic state is supported by the implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExtendedDynamicState3FeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_extended_dynamic_state3
typedef struct VkPhysicalDeviceExtendedDynamicState3FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           extendedDynamicState3TessellationDomainOrigin;
    VkBool32           extendedDynamicState3DepthClampEnable;
    VkBool32           extendedDynamicState3PolygonMode;
    VkBool32           extendedDynamicState3RasterizationSamples;
    VkBool32           extendedDynamicState3SampleMask;
    VkBool32           extendedDynamicState3AlphaToCoverageEnable;
    VkBool32           extendedDynamicState3AlphaToOneEnable;
    VkBool32           extendedDynamicState3LogicOpEnable;
    VkBool32           extendedDynamicState3ColorBlendEnable;
    VkBool32           extendedDynamicState3ColorBlendEquation;
    VkBool32           extendedDynamicState3ColorWriteMask;
    VkBool32           extendedDynamicState3RasterizationStream;
    VkBool32           extendedDynamicState3ConservativeRasterizationMode;
    VkBool32           extendedDynamicState3ExtraPrimitiveOverestimationSize;
    VkBool32           extendedDynamicState3DepthClipEnable;
    VkBool32           extendedDynamicState3SampleLocationsEnable;
    VkBool32           extendedDynamicState3ColorBlendAdvanced;
    VkBool32           extendedDynamicState3ProvokingVertexMode;
    VkBool32           extendedDynamicState3LineRasterizationMode;
    VkBool32           extendedDynamicState3LineStippleEnable;
    VkBool32           extendedDynamicState3DepthClipNegativeOneToOne;
    VkBool32           extendedDynamicState3ViewportWScalingEnable;
    VkBool32           extendedDynamicState3ViewportSwizzle;
    VkBool32           extendedDynamicState3CoverageToColorEnable;
    VkBool32           extendedDynamicState3CoverageToColorLocation;
    VkBool32           extendedDynamicState3CoverageModulationMode;
    VkBool32           extendedDynamicState3CoverageModulationTableEnable;
    VkBool32           extendedDynamicState3CoverageModulationTable;
    VkBool32           extendedDynamicState3CoverageReductionMode;
    VkBool32           extendedDynamicState3RepresentativeFragmentTestEnable;
    VkBool32           extendedDynamicState3ShadingRateImageEnable;
} VkPhysicalDeviceExtendedDynamicState3FeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`extendedDynamicState3TessellationDomainOrigin` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT`
- []()`extendedDynamicState3DepthClampEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT`
- []()`extendedDynamicState3PolygonMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_POLYGON_MODE_EXT`
- []()`extendedDynamicState3RasterizationSamples` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`
- []()`extendedDynamicState3SampleMask` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`
- []()`extendedDynamicState3AlphaToCoverageEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT`
- []()`extendedDynamicState3AlphaToOneEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT`
- []()`extendedDynamicState3LogicOpEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`
- []()`extendedDynamicState3ColorBlendEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`
- []()`extendedDynamicState3ColorBlendEquation` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`
- []()`extendedDynamicState3ColorWriteMask` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`
- []()`extendedDynamicState3RasterizationStream` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT`
- []()`extendedDynamicState3ConservativeRasterizationMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`
- []()`extendedDynamicState3ExtraPrimitiveOverestimationSize` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`
- []()`extendedDynamicState3DepthClipEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT`
- []()`extendedDynamicState3SampleLocationsEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT`
- []()`extendedDynamicState3ColorBlendAdvanced` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`
- []()`extendedDynamicState3ProvokingVertexMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT`
- []()`extendedDynamicState3LineRasterizationMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`
- []()`extendedDynamicState3LineStippleEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT`
- []()`extendedDynamicState3DepthClipNegativeOneToOne` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT`
- []()`extendedDynamicState3ViewportWScalingEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV`
- []()`extendedDynamicState3ViewportSwizzle` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV`
- []()`extendedDynamicState3CoverageToColorEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV`
- []()`extendedDynamicState3CoverageToColorLocation` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV`
- []()`extendedDynamicState3CoverageModulationMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`
- []()`extendedDynamicState3CoverageModulationTableEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV`
- []()`extendedDynamicState3CoverageModulationTable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`
- []()`extendedDynamicState3CoverageReductionMode` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`
- []()`extendedDynamicState3RepresentativeFragmentTestEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV`
- []()`extendedDynamicState3ShadingRateImageEnable` indicates that the implementation supports the following dynamic state:
  
  - `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV`

## [](#_description)Description

If the `VkPhysicalDeviceExtendedDynamicState3FeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceExtendedDynamicState3FeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExtendedDynamicState3FeaturesEXT-sType-sType)VUID-VkPhysicalDeviceExtendedDynamicState3FeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExtendedDynamicState3FeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0