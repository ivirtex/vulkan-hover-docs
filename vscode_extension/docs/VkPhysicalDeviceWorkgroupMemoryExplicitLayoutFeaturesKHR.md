# VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR - Structure describing the workgroup storage explicit layout features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_workgroup_memory_explicit_layout
typedef struct VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           workgroupMemoryExplicitLayout;
    VkBool32           workgroupMemoryExplicitLayoutScalarBlockLayout;
    VkBool32           workgroupMemoryExplicitLayout8BitAccess;
    VkBool32           workgroupMemoryExplicitLayout16BitAccess;
} VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`workgroupMemoryExplicitLayout` indicates whether the implementation supports the SPIR-V `WorkgroupMemoryExplicitLayoutKHR` capability.
- []()`workgroupMemoryExplicitLayoutScalarBlockLayout` indicates whether the implementation supports scalar alignment for laying out Workgroup Blocks.
- []()`workgroupMemoryExplicitLayout8BitAccess` indicates whether objects in the `Workgroup` storage class with the `Block` decoration **can** have 8-bit integer members. If this feature is not enabled, 8-bit integer members **must** not be used in such objects. This also indicates whether shader modules **can** declare the `WorkgroupMemoryExplicitLayout8BitAccessKHR` capability.
- []()`workgroupMemoryExplicitLayout16BitAccess` indicates whether objects in the `Workgroup` storage class with the `Block` decoration **can** have 16-bit integer and 16-bit floating-point members. If this feature is not enabled, 16-bit integer or 16-bit floating-point members **must** not be used in such objects. This also indicates whether shader modules **can** declare the `WorkgroupMemoryExplicitLayout16BitAccessKHR` capability.

## [](#_description)Description

If the `VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_workgroup\_memory\_explicit\_layout](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_workgroup_memory_explicit_layout.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0