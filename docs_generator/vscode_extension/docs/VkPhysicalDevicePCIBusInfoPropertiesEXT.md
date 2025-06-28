# VkPhysicalDevicePCIBusInfoPropertiesEXT(3) Manual Page

## Name

VkPhysicalDevicePCIBusInfoPropertiesEXT - Structure containing PCI bus information of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDevicePCIBusInfoPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_pci_bus_info
typedef struct VkPhysicalDevicePCIBusInfoPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           pciDomain;
    uint32_t           pciBus;
    uint32_t           pciDevice;
    uint32_t           pciFunction;
} VkPhysicalDevicePCIBusInfoPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pciDomain` is the PCI bus domain.
- `pciBus` is the PCI bus identifier.
- `pciDevice` is the PCI device identifier.
- `pciFunction` is the PCI device function identifier.

## [](#_description)Description

If the `VkPhysicalDevicePCIBusInfoPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the PCI bus information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePCIBusInfoPropertiesEXT-sType-sType)VUID-VkPhysicalDevicePCIBusInfoPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_pci\_bus\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pci_bus_info.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePCIBusInfoPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0