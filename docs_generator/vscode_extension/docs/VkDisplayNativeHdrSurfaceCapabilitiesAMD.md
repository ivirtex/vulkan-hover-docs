# VkDisplayNativeHdrSurfaceCapabilitiesAMD(3) Manual Page

## Name

VkDisplayNativeHdrSurfaceCapabilitiesAMD - Structure describing display native HDR specific capabilities of a surface



## [](#_c_specification)C Specification

The `VkDisplayNativeHdrSurfaceCapabilitiesAMD` structure is defined as:

```c++
// Provided by VK_AMD_display_native_hdr
typedef struct VkDisplayNativeHdrSurfaceCapabilitiesAMD {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           localDimmingSupport;
} VkDisplayNativeHdrSurfaceCapabilitiesAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `localDimmingSupport` specifies whether the surface supports local dimming. If this is `VK_TRUE`, [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html) **can** be used to explicitly enable or disable local dimming for the surface. Local dimming may also be overridden by [vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLocalDimmingAMD.html) during the lifetime of the swapchain.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayNativeHdrSurfaceCapabilitiesAMD-sType-sType)VUID-VkDisplayNativeHdrSurfaceCapabilitiesAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD`

## [](#_see_also)See Also

[VK\_AMD\_display\_native\_hdr](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_display_native_hdr.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayNativeHdrSurfaceCapabilitiesAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0