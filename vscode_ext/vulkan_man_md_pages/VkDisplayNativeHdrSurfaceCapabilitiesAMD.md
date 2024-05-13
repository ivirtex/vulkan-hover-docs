# VkDisplayNativeHdrSurfaceCapabilitiesAMD(3) Manual Page

## Name

VkDisplayNativeHdrSurfaceCapabilitiesAMD - Structure describing display
native HDR specific capabilities of a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayNativeHdrSurfaceCapabilitiesAMD` structure is defined as:

``` c
// Provided by VK_AMD_display_native_hdr
typedef struct VkDisplayNativeHdrSurfaceCapabilitiesAMD {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           localDimmingSupport;
} VkDisplayNativeHdrSurfaceCapabilitiesAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `localDimmingSupport` specifies whether the surface supports local
  dimming. If this is `VK_TRUE`,
  [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html)
  **can** be used to explicitly enable or disable local dimming for the
  surface. Local dimming may also be overridden by
  [vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLocalDimmingAMD.html) during the lifetime
  of the swapchain.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayNativeHdrSurfaceCapabilitiesAMD-sType-sType"
  id="VUID-VkDisplayNativeHdrSurfaceCapabilitiesAMD-sType-sType"></a>
  VUID-VkDisplayNativeHdrSurfaceCapabilitiesAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_display_native_hdr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_display_native_hdr.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayNativeHdrSurfaceCapabilitiesAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
