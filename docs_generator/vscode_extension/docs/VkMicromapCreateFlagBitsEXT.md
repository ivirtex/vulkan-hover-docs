# VkMicromapCreateFlagBitsEXT(3) Manual Page

## Name

VkMicromapCreateFlagBitsEXT - Bitmask specifying additional creation parameters for micromap



## [](#_c_specification)C Specification

Bits which **can** be set in [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`createFlags`, specifying additional creation parameters for micromaps, are:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkMicromapCreateFlagBitsEXT {
    VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT = 0x00000001,
} VkMicromapCreateFlagBitsEXT;
```

## [](#_description)Description

- `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT` specifies that the micromap’s address **can** be saved and reused on a subsequent run.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapCreateFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0