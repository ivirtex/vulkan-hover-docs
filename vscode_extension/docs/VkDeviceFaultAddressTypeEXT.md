# VkDeviceFaultAddressTypeEXT(3) Manual Page

## Name

VkDeviceFaultAddressTypeEXT - Page fault access types



## [](#_c_specification)C Specification

Possible values of [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html)::`addressType` are:

```c++
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

## [](#_description)Description

- `VK_DEVICE_FAULT_ADDRESS_TYPE_NONE_EXT` specifies that [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) does not describe a page fault, or an instruction address.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_READ_INVALID_EXT` specifies that [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) describes a page fault triggered by an invalid read operation.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_WRITE_INVALID_EXT` specifies that [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) describes a page fault triggered by an invalid write operation.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_EXECUTE_INVALID_EXT` describes a page fault triggered by an attempt to execute non-executable memory.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_UNKNOWN_EXT` specifies an instruction pointer value at the time the fault occurred. This may or may not be related to a fault.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_INVALID_EXT` specifies an instruction pointer value associated with an invalid instruction fault.
- `VK_DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_FAULT_EXT` specifies an instruction pointer value associated with a fault.

Note

The instruction pointer values recorded may not identify the specific instruction(s) that triggered the fault. The relationship between the instruction pointer reported and triggering instruction will be vendor-specific.

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultAddressTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0