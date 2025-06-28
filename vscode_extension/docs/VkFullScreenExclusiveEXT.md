# VkFullScreenExclusiveEXT(3) Manual Page

## Name

VkFullScreenExclusiveEXT - Hint values an application can specify affecting full-screen transition behavior



## [](#_c_specification)C Specification

Possible values of `VkSurfaceFullScreenExclusiveInfoEXT`::`fullScreenExclusive` are:

```c++
// Provided by VK_EXT_full_screen_exclusive
typedef enum VkFullScreenExclusiveEXT {
    VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT = 0,
    VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT = 1,
    VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT = 2,
    VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT = 3,
} VkFullScreenExclusiveEXT;
```

## [](#_description)Description

- `VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT` specifies that the implementation **should** determine the appropriate full-screen method by whatever means it deems appropriate.
- `VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT` specifies that the implementation **may** use full-screen exclusive mechanisms when available. Such mechanisms **may** result in better performance and/or the availability of different presentation capabilities, but **may** require a more disruptive transition during swapchain initialization, first presentation and/or destruction.
- `VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT` specifies that the implementation **should** avoid using full-screen mechanisms which rely on disruptive transitions.
- `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT` specifies that the application will manage full-screen exclusive mode by using the [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireFullScreenExclusiveModeEXT.html) and [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseFullScreenExclusiveModeEXT.html) commands.

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFullScreenExclusiveEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0