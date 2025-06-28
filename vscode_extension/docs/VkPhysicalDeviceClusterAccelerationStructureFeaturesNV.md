# VkPhysicalDeviceClusterAccelerationStructureFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceClusterAccelerationStructureFeaturesNV - Structure describing the ray tracing cluster geometry feature supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceClusterAccelerationStructureFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkPhysicalDeviceClusterAccelerationStructureFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           clusterAccelerationStructure;
} VkPhysicalDeviceClusterAccelerationStructureFeaturesNV;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`clusterAccelerationStructure` indicates whether the implementation supports the ability to generate and trace cluster acceleration structures.

## [](#_description)Description

If the `VkPhysicalDeviceClusterAccelerationStructureFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceClusterAccelerationStructureFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceClusterAccelerationStructureFeaturesNV-sType-sType)VUID-VkPhysicalDeviceClusterAccelerationStructureFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_ACCELERATION_STRUCTURE_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceClusterAccelerationStructureFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0