# VkPhysicalDeviceConservativeRasterizationPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceConservativeRasterizationPropertiesEXT - Structure
describing conservative raster properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceConservativeRasterizationPropertiesEXT` structure
is defined as:

``` c
// Provided by VK_EXT_conservative_rasterization
typedef struct VkPhysicalDeviceConservativeRasterizationPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    float              primitiveOverestimationSize;
    float              maxExtraPrimitiveOverestimationSize;
    float              extraPrimitiveOverestimationSizeGranularity;
    VkBool32           primitiveUnderestimation;
    VkBool32           conservativePointAndLineRasterization;
    VkBool32           degenerateTrianglesRasterized;
    VkBool32           degenerateLinesRasterized;
    VkBool32           fullyCoveredFragmentShaderInputVariable;
    VkBool32           conservativeRasterizationPostDepthCoverage;
} VkPhysicalDeviceConservativeRasterizationPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-primitiveOverestimationSize"></span>
  `primitiveOverestimationSize` is the size in pixels the generating
  primitive is increased at each of its edges during conservative
  rasterization overestimation mode. Even with a size of 0.0,
  conservative rasterization overestimation rules still apply and if any
  part of the pixel rectangle is covered by the generating primitive,
  fragments are generated for the entire pixel. However implementations
  **may** make the pixel coverage area even more conservative by
  increasing the size of the generating primitive.

- <span id="limits-maxExtraPrimitiveOverestimationSize"></span>
  `maxExtraPrimitiveOverestimationSize` is the maximum size in pixels of
  extra overestimation the implementation supports in the pipeline
  state. A value of 0.0 means the implementation does not support any
  additional overestimation of the generating primitive during
  conservative rasterization. A value above 0.0 allows the application
  to further increase the size of the generating primitive during
  conservative rasterization overestimation.

- <span id="limits-extraPrimitiveOverestimationSizeGranularity"></span>
  `extraPrimitiveOverestimationSizeGranularity` is the granularity of
  extra overestimation that can be specified in the pipeline state
  between 0.0 and `maxExtraPrimitiveOverestimationSize` inclusive. A
  value of 0.0 means the implementation can use the smallest
  representable non-zero value in the screen space pixel fixed-point
  grid.

- <span id="limits-primitiveUnderestimation"></span>
  `primitiveUnderestimation` is `VK_TRUE` if the implementation supports
  the `VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT`
  conservative rasterization mode in addition to
  `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT`. Otherwise the
  implementation only supports
  `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT`.

- <span id="limits-conservativePointAndLineRasterization"></span>
  `conservativePointAndLineRasterization` is `VK_TRUE` if the
  implementation supports conservative rasterization of point and line
  primitives as well as triangle primitives. Otherwise the
  implementation only supports triangle primitives.

- <span id="limits-degenerateTrianglesRasterized"></span>
  `degenerateTrianglesRasterized` is `VK_FALSE` if the implementation
  culls primitives generated from triangles that become zero area after
  they are quantized to the fixed-point rasterization pixel grid.
  `degenerateTrianglesRasterized` is `VK_TRUE` if these primitives are
  not culled and the provoking vertex attributes and depth value are
  used for the fragments. The primitive area calculation is done on the
  primitive generated from the clipped triangle if applicable. Zero area
  primitives are backfacing and the application **can** enable backface
  culling if desired.

- <span id="limits-degenerateLinesRasterized"></span>
  `degenerateLinesRasterized` is `VK_FALSE` if the implementation culls
  lines that become zero length after they are quantized to the
  fixed-point rasterization pixel grid. `degenerateLinesRasterized` is
  `VK_TRUE` if zero length lines are not culled and the provoking vertex
  attributes and depth value are used for the fragments.

- <span id="limits-fullyCoveredFragmentShaderInputVariable"></span>
  `fullyCoveredFragmentShaderInputVariable` is `VK_TRUE` if the
  implementation supports the SPIR-V builtin fragment shader input
  variable `FullyCoveredEXT` specifying that conservative rasterization
  is enabled and the fragment area is fully covered by the generating
  primitive.

- <span id="limits-conservativeRasterizationPostDepthCoverage"></span>
  `conservativeRasterizationPostDepthCoverage` is `VK_TRUE` if the
  implementation supports conservative rasterization with the
  `PostDepthCoverage` execution mode enabled. Otherwise the
  `PostDepthCoverage` execution mode **must** not be used when
  conservative rasterization is enabled.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceConservativeRasterizationPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceConservativeRasterizationPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceConservativeRasterizationPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceConservativeRasterizationPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
