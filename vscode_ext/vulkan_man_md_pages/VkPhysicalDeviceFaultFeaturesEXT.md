# VkPhysicalDeviceFaultFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceFaultFeaturesEXT - Structure indicating support for
device fault reporting



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFaultFeaturesEXT` structure is defined as:

``` c
// Provided by VK_EXT_device_fault
typedef struct VkPhysicalDeviceFaultFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           deviceFault;
    VkBool32           deviceFaultVendorBinary;
} VkPhysicalDeviceFaultFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceFaultFeaturesEXT` structure describe
the following features:

## <a href="#_description" class="anchor"></a>Description

- <span id="features-deviceFault"></span> `deviceFault` indicates that
  the implementation supports the reporting of device fault information.

- <span id="features-deviceFaultVendorBinary"></span>
  `deviceFaultVendorBinary` indicates that the implementation supports
  the generation of vendor-specific binary crash dumps. These may
  provide additional information when imported into vendor-specific
  external tools.

If the `VkPhysicalDeviceFaultFeaturesEXT` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceFaultFeaturesEXT` **can** also be used in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively
enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceFaultFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceFaultFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceFaultFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FAULT_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFaultFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
