# VkSwapchainDisplayNativeHdrCreateInfoAMD(3) Manual Page

## Name

VkSwapchainDisplayNativeHdrCreateInfoAMD - Structure specifying display
native HDR parameters of a newly created swapchain object



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) includes a
`VkSwapchainDisplayNativeHdrCreateInfoAMD` structure, then that
structure includes additional swapchain creation parameters specific to
display native HDR support.

The `VkSwapchainDisplayNativeHdrCreateInfoAMD` structure is defined as:

``` c
// Provided by VK_AMD_display_native_hdr
typedef struct VkSwapchainDisplayNativeHdrCreateInfoAMD {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           localDimmingEnable;
} VkSwapchainDisplayNativeHdrCreateInfoAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `localDimmingEnable` specifies whether local dimming is enabled for
  the swapchain.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) does not
include this structure, the default value for `localDimmingEnable` is
`VK_TRUE`, meaning local dimming is initially enabled for the swapchain.

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-sType-sType"
  id="VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-sType-sType"></a>
  VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD`

Valid Usage

- <a
  href="#VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-localDimmingEnable-04449"
  id="VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-localDimmingEnable-04449"></a>
  VUID-VkSwapchainDisplayNativeHdrCreateInfoAMD-localDimmingEnable-04449  
  It is only valid to set `localDimmingEnable` to `VK_TRUE` if
  [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)::`localDimmingSupport`
  is supported

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_display_native_hdr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_display_native_hdr.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainDisplayNativeHdrCreateInfoAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
