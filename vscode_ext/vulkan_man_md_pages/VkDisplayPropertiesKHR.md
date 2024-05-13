# VkDisplayPropertiesKHR(3) Manual Page

## Name

VkDisplayPropertiesKHR - Structure describing an available display
device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayPropertiesKHR {
    VkDisplayKHR                  display;
    const char*                   displayName;
    VkExtent2D                    physicalDimensions;
    VkExtent2D                    physicalResolution;
    VkSurfaceTransformFlagsKHR    supportedTransforms;
    VkBool32                      planeReorderPossible;
    VkBool32                      persistentContent;
} VkDisplayPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `display` is a handle that is used to refer to the display described
  here. This handle will be valid for the lifetime of the Vulkan
  instance.

- `displayName` is `NULL` or a pointer to a null-terminated UTF-8 string
  containing the name of the display. Generally, this will be the name
  provided by the display’s EDID. If `NULL`, no suitable name is
  available. If not `NULL`, the string pointed to **must** remain
  accessible and unmodified as long as `display` is valid.

- `physicalDimensions` describes the physical width and height of the
  visible portion of the display, in millimeters.

- `physicalResolution` describes the physical, native, or preferred
  resolution of the display.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For devices which have no natural value to return here,
implementations <strong>should</strong> return the maximum resolution
supported.</p></td>
</tr>
</tbody>
</table>

- `supportedTransforms` is a bitmask of
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  describing which transforms are supported by this display.

- `planeReorderPossible` tells whether the planes on this display
  **can** have their z order changed. If this is `VK_TRUE`, the
  application **can** re-arrange the planes on this display in any order
  relative to each other.

- `persistentContent` tells whether the display supports
  self-refresh/internal buffering. If this is true, the application
  **can** submit persistent present operations on swapchains created
  against this display.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Persistent presents <strong>may</strong> have higher latency, and
<strong>may</strong> use less power when the screen content is updated
infrequently, or when only a portion of the screen needs to be updated
in most frames.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayProperties2KHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagsKHR.html),
[vkGetPhysicalDeviceDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
