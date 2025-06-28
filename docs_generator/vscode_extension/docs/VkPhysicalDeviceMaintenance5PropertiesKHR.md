# VkPhysicalDeviceMaintenance5Properties(3) Manual Page

## Name

VkPhysicalDeviceMaintenance5Properties - Structure describing various implementation-defined properties introduced with VK\_KHR\_maintenance5



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance5Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceMaintenance5Properties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           earlyFragmentMultisampleCoverageAfterSampleCounting;
    VkBool32           earlyFragmentSampleMaskTestBeforeSampleCounting;
    VkBool32           depthStencilSwizzleOneSupport;
    VkBool32           polygonModePointSize;
    VkBool32           nonStrictSinglePixelWideLinesUseParallelogram;
    VkBool32           nonStrictWideLinesUseParallelogram;
} VkPhysicalDeviceMaintenance5Properties;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkPhysicalDeviceMaintenance5Properties VkPhysicalDeviceMaintenance5PropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `earlyFragmentMultisampleCoverageAfterSampleCounting` is a boolean value indicating whether the [fragment shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader) and [multisample coverage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-covg) operations are performed after [sample counting](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-samplecount) for [fragment shaders](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader) with `EarlyFragmentTests` execution mode.
- `earlyFragmentSampleMaskTestBeforeSampleCounting` is a boolean value indicating whether the [sample mask test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-samplemask) operation is performed before [sample counting](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-samplecount) for [fragment shaders](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader) using the `EarlyFragmentTests` execution mode.
- `depthStencilSwizzleOneSupport` is a boolean indicating that depth/stencil texturing operations with `VK_COMPONENT_SWIZZLE_ONE` have defined behavior.
- `polygonModePointSize` is a boolean value indicating whether the point size of the final rasterization of polygons with `VK_POLYGON_MODE_POINT` is controlled by `PointSize`.
- `nonStrictSinglePixelWideLinesUseParallelogram` is a boolean value indicating whether non-strict lines with a width of 1.0 are rasterized as parallelograms or using Bresenham’s algorithm.
- `nonStrictWideLinesUseParallelogram` is a boolean value indicating whether non-strict lines with a width greater than 1.0 are rasterized as parallelograms or using Bresenham’s algorithm.

If the `VkPhysicalDeviceMaintenance5Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance5Properties-sType-sType)VUID-VkPhysicalDeviceMaintenance5Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_PROPERTIES`

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance5Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0