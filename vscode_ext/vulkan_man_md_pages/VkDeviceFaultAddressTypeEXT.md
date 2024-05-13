# VkDeviceFaultAddressTypeEXT(3) Manual Page

## Name

VkDeviceFaultAddressTypeEXT - Page fault access types



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)::`addressType`
are:

``` c
// Provided by VK_EXT_device_fault
typedef enum VkDeviceFaultAddressTypeEXT {
    VK_DEVICE_FAULT_ADDRESS_TYPE_NONE_EXT = 0,
    VK_DEVICE_FAULT_ADDRESS_TYPE_READ_INVALID_EXT = 1,
    VK_DEVICE_FAULT_ADDRESS_TYPE_WRITE_INVALID_EXT = 2,
    VK_DEVICE_FAULT_ADDRESS_TYPE_EXECUTE_INVALID_EXT = 3,
    VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_UNKNOWN_EXT = 4,
    VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_INVALID_EXT = 5,
    VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_FAULT_EXT = 6,
} VkDeviceFaultAddressTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_FAULT_ADDRESS_TYPE_NONE_EXT` specifies that
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html) does
  not describe a page fault, or an instruction address.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_READ_INVALID_EXT` specifies that
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)
  describes a page fault triggered by an invalid read operation.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_WRITE_INVALID_EXT` specifies that
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)
  describes a page fault triggered by an invalid write operation.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_EXECUTE_INVALID_EXT` describes a page
  fault triggered by an attempt to execute non-executable memory.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_UNKNOWN_EXT`
  specifies an instruction pointer value at the time the fault occurred.
  This may or may not be related to a fault.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_INVALID_EXT`
  specifies an instruction pointer value associated with an invalid
  instruction fault.

- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_FAULT_EXT` specifies
  an instruction pointer value associated with a fault.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The instruction pointer values recorded may not identify the specific
instruction(s) that triggered the fault. The relationship between the
instruction pointer reported and triggering instruction will be
vendor-specific.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultAddressTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
