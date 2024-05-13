# VkMicromapCreateFlagBitsEXT(3) Manual Page

## Name

VkMicromapCreateFlagBitsEXT - Bitmask specifying additional creation
parameters for micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)::`createFlags`,
specifying additional creation parameters for micromaps, are:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkMicromapCreateFlagBitsEXT {
    VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT = 0x00000001,
} VkMicromapCreateFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT` specifies
  that the micromap’s address **can** be saved and reused on a
  subsequent run.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapCreateFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
