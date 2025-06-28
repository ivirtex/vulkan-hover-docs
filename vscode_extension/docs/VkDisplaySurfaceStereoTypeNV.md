# VkDisplaySurfaceStereoTypeNV(3) Manual Page

## Name

VkDisplaySurfaceStereoTypeNV - 3D Stereo type



## [](#_c_specification)C Specification

Possible values of [VkDisplaySurfaceStereoCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoCreateInfoNV.html)::`stereoType`, specifying the type of 3D stereo presentation the display will be configured for, are:

```c++
// Provided by VK_NV_display_stereo
typedef enum VkDisplaySurfaceStereoTypeNV {
    VK_DISPLAY_SURFACE_STEREO_TYPE_NONE_NV = 0,
    VK_DISPLAY_SURFACE_STEREO_TYPE_ONBOARD_DIN_NV = 1,
    VK_DISPLAY_SURFACE_STEREO_TYPE_HDMI_3D_NV = 2,
    VK_DISPLAY_SURFACE_STEREO_TYPE_INBAND_DISPLAYPORT_NV = 3,
} VkDisplaySurfaceStereoTypeNV;
```

## [](#_description)Description

- `VK_DISPLAY_SURFACE_STEREO_TYPE_NONE_NV` specifies no configuration for stereo presentation. This is the default behavior if [VkDisplaySurfaceStereoCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoCreateInfoNV.html) is not provided.
- `VK_DISPLAY_SURFACE_STEREO_TYPE_ONBOARD_DIN_NV` specifies configuration for glasses that connect via a DIN connector on the back of the graphics card.
- `VK_DISPLAY_SURFACE_STEREO_TYPE_HDMI_3D_NV` specifies configuration for HDMI 3D compatible display devices with their own stereo emitters. This is also known as HDMI Frame Packed Stereo, where the left and right eye images are stacked into a single frame with a doubled pixel clock and refresh rate.
- `VK_DISPLAY_SURFACE_STEREO_TYPE_INBAND_DISPLAYPORT_NV` specifies configuration for DisplayPort display devices with in-band stereo signaling and emitters.

## [](#_see_also)See Also

[VK\_NV\_display\_stereo](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_display_stereo.html), [VkDisplaySurfaceStereoCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoCreateInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplaySurfaceStereoTypeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0