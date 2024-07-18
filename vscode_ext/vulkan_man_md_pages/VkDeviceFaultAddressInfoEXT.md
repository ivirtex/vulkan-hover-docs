# VkDeviceFaultAddressInfoEXT(3) Manual Page

## Name

VkDeviceFaultAddressInfoEXT - Structure specifying GPU virtual address
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceFaultAddressInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultAddressInfoEXT {
    VkDeviceFaultAddressTypeEXT    addressType;
    VkDeviceAddress                reportedAddress;
    VkDeviceSize                   addressPrecision;
} VkDeviceFaultAddressInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `addressType` is either the type of memory operation that triggered a
  page fault, or the type of association between an instruction pointer
  and a fault.

- `reportedAddress` is the GPU virtual address recorded by the device.

- `addressPrecision` is a power of two value that specifies how
  precisely the device can report the address.

## <a href="#_description" class="anchor"></a>Description

The combination of `reportedAddress` and `addressPrecision` allow the
possible range of addresses to be calculated, such that:

``` c
lower_address = (pInfo->reportedAddress & ~(pInfo->addressPrecision-1))
upper_address = (pInfo->reportedAddress |  (pInfo->addressPrecision-1))
```

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is valid for the <code>reportedAddress</code> to contain a more
precise address than indicated by <code>addressPrecision</code>. In this
case, the value of <code>reportedAddress</code> should be treated as an
additional hint as to the value of the address that triggered the page
fault, or to the value of an instruction pointer.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceFaultAddressInfoEXT-addressType-parameter"
  id="VUID-VkDeviceFaultAddressInfoEXT-addressType-parameter"></a>
  VUID-VkDeviceFaultAddressInfoEXT-addressType-parameter  
  `addressType` **must** be a valid
  [VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressTypeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressTypeEXT.html),
[VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultAddressInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
