# VkDisplayPlaneInfo2KHR(3) Manual Page

## Name

VkDisplayPlaneInfo2KHR - Structure defining the intended configuration
of a display plane



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPlaneInfo2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayPlaneInfo2KHR {
    VkStructureType     sType;
    const void*         pNext;
    VkDisplayModeKHR    mode;
    uint32_t            planeIndex;
} VkDisplayPlaneInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mode` is the display mode the application intends to program when
  using the specified plane.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This parameter also implicitly specifies a display.</p></td>
</tr>
</tbody>
</table>

- `planeIndex` is the plane which the application intends to use with
  the display.

The members of `VkDisplayPlaneInfo2KHR` correspond to the arguments to
[vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html),
with `sType` and `pNext` added for extensibility.

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayPlaneInfo2KHR-sType-sType"
  id="VUID-VkDisplayPlaneInfo2KHR-sType-sType"></a>
  VUID-VkDisplayPlaneInfo2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR`

- <a href="#VUID-VkDisplayPlaneInfo2KHR-pNext-pNext"
  id="VUID-VkDisplayPlaneInfo2KHR-pNext-pNext"></a>
  VUID-VkDisplayPlaneInfo2KHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDisplayPlaneInfo2KHR-mode-parameter"
  id="VUID-VkDisplayPlaneInfo2KHR-mode-parameter"></a>
  VUID-VkDisplayPlaneInfo2KHR-mode-parameter  
  `mode` **must** be a valid [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html)
  handle

Host Synchronization

- Host access to `mode` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilities2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPlaneInfo2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
