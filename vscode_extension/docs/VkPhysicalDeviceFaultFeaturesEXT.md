# VkPhysicalDeviceFaultFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceFaultFeaturesEXT - Structure indicating support for device fault reporting



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFaultFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkPhysicalDeviceFaultFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           deviceFault;
    VkBool32           deviceFaultVendorBinary;
} VkPhysicalDeviceFaultFeaturesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceFaultFeaturesEXT` structure describe the following features:

## [](#_description)Description

- []()`deviceFault` indicates that the implementation supports the reporting of device fault information.
- []()`deviceFaultVendorBinary` indicates that the implementation supports the generation of vendor-specific binary crash dumps. These may provide additional information when imported into vendor-specific external tools.

If the `VkPhysicalDeviceFaultFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceFaultFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFaultFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceFaultFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FAULT_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFaultFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0