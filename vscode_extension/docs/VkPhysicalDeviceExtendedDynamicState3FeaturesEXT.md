# VkPhysicalDeviceExtendedDynamicState3FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceExtendedDynamicState3FeaturesEXT - Structure describing
what extended dynamic state is supported by the implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExtendedDynamicState3FeaturesEXT` structure is
defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-extendedDynamicState3TessellationDomainOrigin"></span>
  `extendedDynamicState3TessellationDomainOrigin` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT`

- <span id="features-extendedDynamicState3DepthClampEnable"></span>
  `extendedDynamicState3DepthClampEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT`

- <span id="features-extendedDynamicState3PolygonMode"></span>
  `extendedDynamicState3PolygonMode` indicates that the implementation
  supports the following dynamic state:

  - `VK_DYNAMIC_STATE_POLYGON_MODE_EXT`

- <span id="features-extendedDynamicState3RasterizationSamples"></span>
  `extendedDynamicState3RasterizationSamples` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`

- <span id="features-extendedDynamicState3SampleMask"></span>
  `extendedDynamicState3SampleMask` indicates that the implementation
  supports the following dynamic state:

  - `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`

- <span id="features-extendedDynamicState3AlphaToCoverageEnable"></span>
  `extendedDynamicState3AlphaToCoverageEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT`

- <span id="features-extendedDynamicState3AlphaToOneEnable"></span>
  `extendedDynamicState3AlphaToOneEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT`

- <span id="features-extendedDynamicState3LogicOpEnable"></span>
  `extendedDynamicState3LogicOpEnable` indicates that the implementation
  supports the following dynamic state:

  - `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`

- <span id="features-extendedDynamicState3ColorBlendEnable"></span>
  `extendedDynamicState3ColorBlendEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`

- <span id="features-extendedDynamicState3ColorBlendEquation"></span>
  `extendedDynamicState3ColorBlendEquation` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`

- <span id="features-extendedDynamicState3ColorWriteMask"></span>
  `extendedDynamicState3ColorWriteMask` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`

- <span id="features-extendedDynamicState3RasterizationStream"></span>
  `extendedDynamicState3RasterizationStream` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT`

- <span id="features-extendedDynamicState3ConservativeRasterizationMode"></span>
  `extendedDynamicState3ConservativeRasterizationMode` indicates that
  the implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`

- <span id="features-extendedDynamicState3ExtraPrimitiveOverestimationSize"></span>
  `extendedDynamicState3ExtraPrimitiveOverestimationSize` indicates that
  the implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`

- <span id="features-extendedDynamicState3DepthClipEnable"></span>
  `extendedDynamicState3DepthClipEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT`

- <span id="features-extendedDynamicState3SampleLocationsEnable"></span>
  `extendedDynamicState3SampleLocationsEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT`

- <span id="features-extendedDynamicState3ColorBlendAdvanced"></span>
  `extendedDynamicState3ColorBlendAdvanced` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`

- <span id="features-extendedDynamicState3ProvokingVertexMode"></span>
  `extendedDynamicState3ProvokingVertexMode` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT`

- <span id="features-extendedDynamicState3LineRasterizationMode"></span>
  `extendedDynamicState3LineRasterizationMode` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`

- <span id="features-extendedDynamicState3LineStippleEnable"></span>
  `extendedDynamicState3LineStippleEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT`

- <span id="features-extendedDynamicState3DepthClipNegativeOneToOne"></span>
  `extendedDynamicState3DepthClipNegativeOneToOne` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT`

- <span id="features-extendedDynamicState3ViewportWScalingEnable"></span>
  `extendedDynamicState3ViewportWScalingEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV`

- <span id="features-extendedDynamicState3ViewportSwizzle"></span>
  `extendedDynamicState3ViewportSwizzle` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV`

- <span id="features-extendedDynamicState3CoverageToColorEnable"></span>
  `extendedDynamicState3CoverageToColorEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV`

- <span id="features-extendedDynamicState3CoverageToColorLocation"></span>
  `extendedDynamicState3CoverageToColorLocation` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV`

- <span id="features-extendedDynamicState3CoverageModulationMode"></span>
  `extendedDynamicState3CoverageModulationMode` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`

- <span id="features-extendedDynamicState3CoverageModulationTableEnable"></span>
  `extendedDynamicState3CoverageModulationTableEnable` indicates that
  the implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV`

- <span id="features-extendedDynamicState3CoverageModulationTable"></span>
  `extendedDynamicState3CoverageModulationTable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`

- <span id="features-extendedDynamicState3CoverageReductionMode"></span>
  `extendedDynamicState3CoverageReductionMode` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`

- <span id="features-extendedDynamicState3RepresentativeFragmentTestEnable"></span>
  `extendedDynamicState3RepresentativeFragmentTestEnable` indicates that
  the implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV`

- <span id="features-extendedDynamicState3ShadingRateImageEnable"></span>
  `extendedDynamicState3ShadingRateImageEnable` indicates that the
  implementation supports the following dynamic state:

  - `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExtendedDynamicState3FeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceExtendedDynamicState3FeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExtendedDynamicState3FeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceExtendedDynamicState3FeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceExtendedDynamicState3FeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExtendedDynamicState3FeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
