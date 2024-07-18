# VkDeviceFaultInfoEXT(3) Manual Page

## Name

VkDeviceFaultInfoEXT - Structure specifying device fault information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceFaultInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultInfoEXT {
    VkStructureType                 sType;
    void*                           pNext;
    char                            description[VK_MAX_DESCRIPTION_SIZE];
    VkDeviceFaultAddressInfoEXT*    pAddressInfos;
    VkDeviceFaultVendorInfoEXT*     pVendorInfos;
    void*                           pVendorBinaryData;
} VkDeviceFaultInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which is a human readable
  description of the fault.

- `pAddressInfos` is `NULL` or a pointer to an array of
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)
  structures describing either memory accesses which **may** have caused
  a page fault, or describing active instruction pointers at the time of
  the fault. If not `NULL`, each element of `pAddressInfos` describes
  the a bounded region of GPU virtual address space containing either
  the GPU virtual address accessed, or the value of an active
  instruction pointer.

- `pVendorInfos` is `NULL` or a pointer to an array of
  [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorInfoEXT.html)
  structures describing vendor-specific fault information.

- `pVendorBinaryData` is `NULL` or a pointer to `vendorBinarySize`
  number of bytes of data, which will be populated with a
  vendor-specific binary crash dump, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vendor-binary-crash-dumps"
  target="_blank" rel="noopener">Vendor Binary Crash Dumps</a>.

## <a href="#_description" class="anchor"></a>Description

An implementation **should** populate as many members of
[VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html) as possible, given the
information available at the time of the fault and the constraints of
the implementation itself.

Due to hardware limitations, `pAddressInfos` describes ranges of GPU
virtual address space, rather than precise addresses. The precise memory
address accessed or the precise value of the instruction pointer
**must** lie within the region described.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Each element of <code>pAddressInfos</code> describes either:</p>
<ul>
<li><p>A memory access which may have triggered a page fault and may
have contributed to device loss</p></li>
<li><p>The value of an active instruction pointer at the time a fault
occurred. This value may be indicative of the active pipeline or shader
at the time of device loss</p></li>
</ul>
<p>Comparison of the GPU virtual addresses described by
<code>pAddressInfos</code> to GPU virtual address ranges reported by the
<a
href="VK_EXT_device_address_binding_report.html"><code>VK_EXT_device_address_binding_report</code></a>
extension may allow applications to correlate between these addresses
and Vulkan objects. Applications should be aware that these addresses
may also correspond to resources internal to an implementation, which
will not be reported via the <a
href="VK_EXT_device_address_binding_report.html"><code>VK_EXT_device_address_binding_report</code></a>
extension.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceFaultInfoEXT-sType-sType"
  id="VUID-VkDeviceFaultInfoEXT-sType-sType"></a>
  VUID-VkDeviceFaultInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_FAULT_INFO_EXT`

- <a href="#VUID-VkDeviceFaultInfoEXT-pNext-pNext"
  id="VUID-VkDeviceFaultInfoEXT-pNext-pNext"></a>
  VUID-VkDeviceFaultInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html),
[VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorInfoEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
