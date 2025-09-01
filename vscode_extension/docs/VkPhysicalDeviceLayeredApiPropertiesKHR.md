# VkPhysicalDeviceLayeredApiPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiPropertiesKHR - Structure describing a single layered implementation underneath the Vulkan physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceLayeredApiPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceLayeredApiPropertiesKHR {
    VkStructureType                  sType;
    void*                            pNext;
    uint32_t                         vendorID;
    uint32_t                         deviceID;
    VkPhysicalDeviceLayeredApiKHR    layeredAPI;
    char                             deviceName[VK_MAX_PHYSICAL_DEVICE_NAME_SIZE];
} VkPhysicalDeviceLayeredApiPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vendorID` is a unique identifier for the vendor of the layered implementation.
- `deviceID` is a unique identifier for the layered implementation among devices available from the vendor.
- `layeredAPI` is a [VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiKHR.html) specifying the API implemented by the layered implementation.
- `deviceName` is an array of `VK_MAX_PHYSICAL_DEVICE_NAME_SIZE` `char` containing a null-terminated UTF-8 string which is the name of the device.

## [](#_description)Description

If `layeredAPI` is `VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR`, additional Vulkan-specific information can be queried by including the [VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html) structure in the `pNext` chain. Otherwise if such a structure is included in the `pNext` chain, it is ignored.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-sType)VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_KHR`
- [](#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-pNext-pNext)VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html)
- [](#VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-unique)VUID-VkPhysicalDeviceLayeredApiPropertiesKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html), [VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiKHR.html), [VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLayeredApiPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0