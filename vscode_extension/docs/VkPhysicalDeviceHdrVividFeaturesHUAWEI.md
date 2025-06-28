# VkPhysicalDeviceHdrVividFeaturesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceHdrVividFeaturesHUAWEI - Structure describing whether HDR Vivid metadata is supported



## [](#_c_specification)C Specification

The `VkPhysicalDeviceHdrVividFeaturesHUAWEI` structure is defined as:

```c++
// Provided by VK_HUAWEI_hdr_vivid
typedef struct VkPhysicalDeviceHdrVividFeaturesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hdrVivid;
} VkPhysicalDeviceHdrVividFeaturesHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`hdrVivid` specifies whether HDR Vivid metadata is supported.

## [](#_description)Description

If the `VkPhysicalDeviceHdrVividFeaturesHUAWEI` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceHdrVividFeaturesHUAWEI`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceHdrVividFeaturesHUAWEI-sType-sType)VUID-VkPhysicalDeviceHdrVividFeaturesHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HDR_VIVID_FEATURES_HUAWEI`

## [](#_see_also)See Also

[VK\_HUAWEI\_hdr\_vivid](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_hdr_vivid.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceHdrVividFeaturesHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0