# VkDeviceOrHostAddressConstAMDX(3) Manual Page

## Name

VkDeviceOrHostAddressConstAMDX - Union specifying a const device or host address



## [](#_c_specification)C Specification

The `VkDeviceOrHostAddressConstAMDX` union is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef union VkDeviceOrHostAddressConstAMDX {
    VkDeviceAddress    deviceAddress;
    const void*        hostAddress;
} VkDeviceOrHostAddressConstAMDX;
```

## [](#_members)Members

- `deviceAddress` is a buffer device address as returned by the [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressKHR.html) command.
- `hostAddress` is a const host memory address.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html), [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceOrHostAddressConstAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0