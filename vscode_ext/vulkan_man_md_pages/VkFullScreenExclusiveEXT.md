# VkFullScreenExclusiveEXT(3) Manual Page

## Name

VkFullScreenExclusiveEXT - Hint values an application can specify
affecting full-screen transition behavior



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
`VkSurfaceFullScreenExclusiveInfoEXT`::`fullScreenExclusive` are:

``` c
// Provided by VK_EXT_full_screen_exclusive
typedef enum VkFullScreenExclusiveEXT {
    VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT = 0,
    VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT = 1,
    VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT = 2,
    VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT = 3,
} VkFullScreenExclusiveEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT` indicates the implementation
  **should** determine the appropriate full-screen method by whatever
  means it deems appropriate.

- `VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT` indicates the implementation
  **may** use full-screen exclusive mechanisms when available. Such
  mechanisms **may** result in better performance and/or the
  availability of different presentation capabilities, but **may**
  require a more disruptive transition during swapchain initialization,
  first presentation and/or destruction.

- `VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT` indicates the implementation
  **should** avoid using full-screen mechanisms which rely on disruptive
  transitions.

- `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT` indicates the
  application will manage full-screen exclusive mode by using the
  [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html)
  and
  [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html)
  commands.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFullScreenExclusiveEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
