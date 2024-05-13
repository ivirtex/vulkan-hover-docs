# VkPhysicalDevicePCIBusInfoPropertiesEXT(3) Manual Page

## Name

VkPhysicalDevicePCIBusInfoPropertiesEXT - Structure containing PCI bus
information of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePCIBusInfoPropertiesEXT` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pciDomain` is the PCI bus domain.

- `pciBus` is the PCI bus identifier.

- `pciDevice` is the PCI device identifier.

- `pciFunction` is the PCI device function identifier.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePCIBusInfoPropertiesEXT` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These are properties of the PCI bus information of a physical device.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDevicePCIBusInfoPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePCIBusInfoPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePCIBusInfoPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pci_bus_info](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pci_bus_info.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePCIBusInfoPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
