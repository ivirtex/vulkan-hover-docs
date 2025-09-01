# VkPhysicalDevicePortabilitySubsetPropertiesKHR(3) Manual Page

## Name

VkPhysicalDevicePortabilitySubsetPropertiesKHR - Structure describing additional properties supported by a portable implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePortabilitySubsetPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_portability_subset
typedef struct VkPhysicalDevicePortabilitySubsetPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           minVertexInputBindingStrideAlignment;
} VkPhysicalDevicePortabilitySubsetPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`minVertexInputBindingStrideAlignment` indicates the minimum alignment for vertex input strides. [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html)::`stride` **must** be a multiple of, and at least as large as, this value. The value **must** be a power of two.

## [](#_description)Description

If the `VkPhysicalDevicePortabilitySubsetPropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePortabilitySubsetPropertiesKHR-sType-sType)VUID-VkPhysicalDevicePortabilitySubsetPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_portability\_subset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_portability_subset.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePortabilitySubsetPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0