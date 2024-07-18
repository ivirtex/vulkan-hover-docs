# VkDisplayModeParametersKHR(3) Manual Page

## Name

VkDisplayModeParametersKHR - Structure describing display parameters
associated with a display mode



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayModeParametersKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayModeParametersKHR {
    VkExtent2D    visibleRegion;
    uint32_t      refreshRate;
} VkDisplayModeParametersKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `visibleRegion` is the 2D extents of the visible region.

- `refreshRate` is a `uint32_t` that is the number of times the display
  is refreshed each second multiplied by 1000.

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
<p>For example, a 60Hz display mode would report a
<code>refreshRate</code> of 60,000.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkDisplayModeParametersKHR-width-01990"
  id="VUID-VkDisplayModeParametersKHR-width-01990"></a>
  VUID-VkDisplayModeParametersKHR-width-01990  
  The `width` member of `visibleRegion` **must** be greater than `0`

- <a href="#VUID-VkDisplayModeParametersKHR-height-01991"
  id="VUID-VkDisplayModeParametersKHR-height-01991"></a>
  VUID-VkDisplayModeParametersKHR-height-01991  
  The `height` member of `visibleRegion` **must** be greater than `0`

- <a href="#VUID-VkDisplayModeParametersKHR-refreshRate-01992"
  id="VUID-VkDisplayModeParametersKHR-refreshRate-01992"></a>
  VUID-VkDisplayModeParametersKHR-refreshRate-01992  
  `refreshRate` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateInfoKHR.html),
[VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModePropertiesKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayModeParametersKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
