# VkPhysicalDeviceLayeredDriverPropertiesMSFT(3) Manual Page

## Name

VkPhysicalDeviceLayeredDriverPropertiesMSFT - Structure containing
information about driver layering for a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLayeredDriverPropertiesMSFT` structure is defined
as:

``` c
// Provided by VK_MSFT_layered_driver
typedef struct VkPhysicalDeviceLayeredDriverPropertiesMSFT {
    VkStructureType                     sType;
    void*                               pNext;
    VkLayeredDriverUnderlyingApiMSFT    underlyingAPI;
} VkPhysicalDeviceLayeredDriverPropertiesMSFT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `underlyingAPI` is a
  [VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayeredDriverUnderlyingApiMSFT.html)
  value indicating which underlying API is used to implement the layered
  driver, or `VK_LAYERED_DRIVER_UNDERLYING_API_NONE_MSFT` if the driver
  is not layered.

## <a href="#_description" class="anchor"></a>Description

These are properties of the driver layering information of a physical
device.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceLayeredDriverPropertiesMSFT-sType-sType"
  id="VUID-VkPhysicalDeviceLayeredDriverPropertiesMSFT-sType-sType"></a>
  VUID-VkPhysicalDeviceLayeredDriverPropertiesMSFT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_DRIVER_PROPERTIES_MSFT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MSFT_layered_driver](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MSFT_layered_driver.html),
[VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayeredDriverUnderlyingApiMSFT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLayeredDriverPropertiesMSFT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
