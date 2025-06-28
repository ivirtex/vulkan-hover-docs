# VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV - Structure describing the ray tracing linear swept spheres features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_linear_swept_spheres
typedef struct VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           spheres;
    VkBool32           linearSweptSpheres;
} VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`spheres` indicates whether the implementation supports sphere primitives in ray tracing.
- []()`linearSweptSpheres` indicates whether the implementation supports linear swept sphere primitives in ray tracing.

## [](#_description)Description

If the `VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV-sType-sType)VUID-VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_LINEAR_SWEPT_SPHERES_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_linear_swept_spheres.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0