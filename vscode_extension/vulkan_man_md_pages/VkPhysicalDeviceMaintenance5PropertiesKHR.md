# VkPhysicalDeviceMaintenance5PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance5PropertiesKHR - Structure describing various
implementation-defined properties introduced with VK_KHR_maintenance5



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance5PropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkPhysicalDeviceMaintenance5PropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           earlyFragmentMultisampleCoverageAfterSampleCounting;
    VkBool32           earlyFragmentSampleMaskTestBeforeSampleCounting;
    VkBool32           depthStencilSwizzleOneSupport;
    VkBool32           polygonModePointSize;
    VkBool32           nonStrictSinglePixelWideLinesUseParallelogram;
    VkBool32           nonStrictWideLinesUseParallelogram;
} VkPhysicalDeviceMaintenance5PropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `earlyFragmentMultisampleCoverageAfterSampleCounting` is a boolean
  value indicating whether the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader"
  target="_blank" rel="noopener">fragment shading</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-covg"
  target="_blank" rel="noopener">multisample coverage</a> operations are
  performed after <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-samplecount"
  target="_blank" rel="noopener">sample counting</a> for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader"
  target="_blank" rel="noopener">fragment shaders</a> with
  `EarlyFragmentTests` execution mode.

- `earlyFragmentSampleMaskTestBeforeSampleCounting` is a boolean value
  indicating whether the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-samplemask"
  target="_blank" rel="noopener">sample mask test</a> operation is
  performed before <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-samplecount"
  target="_blank" rel="noopener">sample counting</a> for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader"
  target="_blank" rel="noopener">fragment shaders</a> using the
  `EarlyFragmentTests` execution mode.

- `depthStencilSwizzleOneSupport` is a boolean indicating that
  depth/stencil texturing operations with `VK_COMPONENT_SWIZZLE_ONE`
  have defined behavior.

- `polygonModePointSize` is a boolean value indicating whether the point
  size of the final rasterization of polygons with
  `VK_POLYGON_MODE_POINT` is controlled by `PointSize`.

- `nonStrictSinglePixelWideLinesUseParallelogram` is a boolean value
  indicating whether non-strict lines with a width of 1.0 are rasterized
  as parallelograms or using Bresenham’s algorithm.

- `nonStrictWideLinesUseParallelogram` is a boolean value indicating
  whether non-strict lines with a width greater than 1.0 are rasterized
  as parallelograms or using Bresenham’s algorithm.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance5PropertiesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance5PropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance5PropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance5PropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance5PropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
