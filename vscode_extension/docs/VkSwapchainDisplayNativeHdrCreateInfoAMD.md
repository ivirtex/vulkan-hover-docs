# VkSwapchainDisplayNativeHdrCreateInfoAMD(3) Manual Page

## Name

VkSwapchainDisplayNativeHdrCreateInfoAMD - Structure specifying display native HDR parameters of a newly created swapchain object



## [](#_c_specification)C Specification

If the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) includes a `VkSwapchainDisplayNativeHdrCreateInfoAMD` structure, then that structure includes additional swapchain creation parameters specific to display native HDR support.

The `VkSwapchainDisplayNativeHdrCreateInfoAMD` structure is defined as:

```c++
// Provided by VK_AMD_display_native_hdr
typedef struct VkSwapchainDisplayNativeHdrCreateInfoAMD {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           localDimmingEnable;
} VkSwapchainDisplayNativeHdrCreateInfoAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `localDimmingEnable` specifies whether local dimming is enabled for the swapchain.

## [](#_description)Description

If the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) does not include this structure, the default value for `localDimmingEnable` is `VK_TRUE`, meaning local dimming is initially enabled for the swapchain.

Valid Usage (Implicit)

- [](#VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-sType-sType)VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD`

Valid Usage

- [](#VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-localDimmingEnable-04449)VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-localDimmingEnable-04449  
  It is only valid to set `localDimmingEnable` to `VK_TRUE` if [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)::`localDimmingSupport` is supported

## [](#_see_also)See Also

[VK\_AMD\_display\_native\_hdr](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_display_native_hdr.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainDisplayNativeHdrCreateInfoAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0