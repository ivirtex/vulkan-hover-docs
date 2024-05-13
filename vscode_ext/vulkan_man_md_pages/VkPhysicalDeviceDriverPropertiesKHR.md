# VkPhysicalDeviceDriverProperties(3) Manual Page

## Name

VkPhysicalDeviceDriverProperties - Structure containing driver
identification information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDriverProperties` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_driver_properties
typedef VkPhysicalDeviceDriverProperties VkPhysicalDeviceDriverPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- `driverID` is a unique identifier for the driver of the physical
  device.

- `driverName` is an array of `VK_MAX_DRIVER_NAME_SIZE` `char`
  containing a null-terminated UTF-8 string which is the name of the
  driver.

- `driverInfo` is an array of `VK_MAX_DRIVER_INFO_SIZE` `char`
  containing a null-terminated UTF-8 string with additional information
  about the driver.

- `conformanceVersion` is the latest version of the Vulkan conformance
  test that the implementor has successfully tested this driver against
  prior to release (see
  [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConformanceVersion.html)).

If the `VkPhysicalDeviceDriverProperties` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the driver corresponding to a physical device.

`driverID` **must** be immutable for a given driver across instances,
processes, driver versions, and system reboots.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDriverProperties-sType-sType"
  id="VUID-VkPhysicalDeviceDriverProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceDriverProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_driver_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_driver_properties.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkConformanceVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConformanceVersion.html),
[VkDriverId](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDriverProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
