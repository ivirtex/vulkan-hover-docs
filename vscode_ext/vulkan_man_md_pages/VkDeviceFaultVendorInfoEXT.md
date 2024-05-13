# VkDeviceFaultVendorInfoEXT(3) Manual Page

## Name

VkDeviceFaultVendorInfoEXT - Structure specifying vendor-specific fault
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceFaultVendorInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultVendorInfoEXT {
    char        description[VK_MAX_DESCRIPTION_SIZE];
    uint64_t    vendorFaultCode;
    uint64_t    vendorFaultData;
} VkDeviceFaultVendorInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which is a human readable
  description of the fault.

- `vendorFaultCode` is the vendor-specific fault code for this fault.

- `vendorFaultData` is the vendor-specific fault data associated with
  this fault.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultVendorInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
