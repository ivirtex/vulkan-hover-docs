# VkDeviceFaultVendorBinaryHeaderVersionEXT(3) Manual Page

## Name

VkDeviceFaultVendorBinaryHeaderVersionEXT - Encode vendor binary crash
dump version



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `headerVersion` value of the crash dump header
are:

``` c
// Provided by VK_EXT_device_fault
typedef enum VkDeviceFaultVendorBinaryHeaderVersionEXT {
    VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT = 1,
} VkDeviceFaultVendorBinaryHeaderVersionEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT` specifies
  version one of the binary crash dump header.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceFaultVendorBinaryHeaderVersionOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionOneEXT.html),
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultVendorBinaryHeaderVersionEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
