# VkPhysicalDeviceDriverProperties(3) Manual Page

## Name

VkPhysicalDeviceDriverProperties - Structure containing driver identification information



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDriverProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceDriverProperties {
    VkStructureType         sType;
    void*                   pNext;
    VkDriverId              driverID;
    char                    driverName[VK_MAX_DRIVER_NAME_SIZE];
    char                    driverInfo[VK_MAX_DRIVER_INFO_SIZE];
    VkConformanceVersion    conformanceVersion;
} VkPhysicalDeviceDriverProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_driver_properties
typedef VkPhysicalDeviceDriverProperties VkPhysicalDeviceDriverPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `driverID` is a unique identifier for the driver of the physical device.
- `driverName` is an array of `VK_MAX_DRIVER_NAME_SIZE` `char` containing a null-terminated UTF-8 string which is the name of the driver.
- `driverInfo` is an array of `VK_MAX_DRIVER_INFO_SIZE` `char` containing a null-terminated UTF-8 string with additional information about the driver.
- `conformanceVersion` is the latest version of the Vulkan conformance test that the implementor has successfully tested this driver against prior to release (see [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersion.html)).

If the `VkPhysicalDeviceDriverProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the driver corresponding to a physical device.

`driverID` **must** be immutable for a given driver across instances, processes, driver versions, and system reboots.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDriverProperties-sType-sType)VUID-VkPhysicalDeviceDriverProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES`

## [](#_see_also)See Also

[VK\_KHR\_driver\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_driver_properties.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersion.html), [VkDriverId](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDriverId.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDriverProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0