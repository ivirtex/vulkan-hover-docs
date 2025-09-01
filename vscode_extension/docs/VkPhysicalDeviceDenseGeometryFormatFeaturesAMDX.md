# VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX(3) Manual Page

## Name

VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX - Structure describing dense geometry format features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_dense_geometry_format
typedef struct VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           denseGeometryFormat;
} VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`denseGeometryFormat` specifies whether the implementation supports DGF1 compressed geometry data.

## [](#_description)Description

If the `VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX-sType-sType)VUID-VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DENSE_GEOMETRY_FORMAT_FEATURES_AMDX`

## [](#_see_also)See Also

[VK\_AMDX\_dense\_geometry\_format](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_dense_geometry_format.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0