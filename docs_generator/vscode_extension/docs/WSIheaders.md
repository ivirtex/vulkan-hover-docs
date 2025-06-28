# WSIheaders(3) Manual Page

## Name

WSIheaders - Control inclusion of window system interface extensions



## [](#_description)Description

To use a Vulkan extension supporting a platform-specific window system, header files for that window system **must** be included at compile time, or platform-specific types **must** be forward-declared. The Vulkan header files are unable to determine whether or not an external header is available at compile time, so platform-specific extensions are provided in separate headers from the core API and platform-independent extensions, allowing applications to decide which ones they need to be defined and how the external headers are included.

Extensions dependent on particular sets of platform headers, or that forward-declare platform-specific types, are declared in a header named for that platform. Before including these platform-specific Vulkan headers, applications **must** include both `vulkan_core.h` and any external native headers the platform extensions depend on.

As a convenience for applications that do not need the flexibility of separate platform-specific Vulkan headers, `vulkan.h` includes `vulkan_core.h`, and then conditionally includes platform-specific Vulkan headers and the external headers they depend on. Applications control which platform-specific headers are included by #defining macros before including `vulkan.h`.

The correspondence between platform-specific extensions, external headers they require, the platform-specific header which declares them, and the preprocessor macros which enable inclusion by `vulkan.h` are shown in the [following table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-wsi-header-table).

Table 1. Window System Extensions and Headers      Extension Name Window System Name Platform-specific Header Required External Headers Controlling `vulkan.h` Macro

`VK_KHR_android_surface`

Android

`vulkan_android.h`

None

`VK_USE_PLATFORM_ANDROID_KHR`

`VK_KHR_wayland_surface`

Wayland

`vulkan_wayland.h`

`<wayland-client.h>`

`VK_USE_PLATFORM_WAYLAND_KHR`

`VK_KHR_win32_surface`, `VK_KHR_external_memory_win32`, `VK_KHR_win32_keyed_mutex`, `VK_KHR_external_semaphore_win32`, `VK_KHR_external_fence_win32`, `VK_NV_external_memory_win32`, `VK_NV_win32_keyed_mutex`

Microsoft Windows

`vulkan_win32.h`

`<windows.h>`

`VK_USE_PLATFORM_WIN32_KHR`

`VK_KHR_xcb_surface`

X11 Xcb

`vulkan_xcb.h`

`<xcb/xcb.h>`

`VK_USE_PLATFORM_XCB_KHR`

`VK_KHR_xlib_surface`

X11 Xlib

`vulkan_xlib.h`

`<X11/Xlib.h>`

`VK_USE_PLATFORM_XLIB_KHR`

`VK_EXT_directfb_surface`

DirectFB

`vulkan_directfb.h`

`<directfb/directfb.h>`

`VK_USE_PLATFORM_DIRECTFB_EXT`

`VK_EXT_acquire_xlib_display`

X11 XRAndR

`vulkan_xlib_xrandr.h`

`<X11/Xlib.h>`, `<X11/extensions/Xrandr.h>`

`VK_USE_PLATFORM_XLIB_XRANDR_EXT`

`VK_GGP_stream_descriptor_surface`, `VK_GGP_frame_token`

Google Games Platform

`vulkan_ggp.h`

&lt;ggp\_c/vulkan\_types.h&gt;

`VK_USE_PLATFORM_GGP`

`VK_MVK_ios_surface`

iOS

`vulkan_ios.h`

None

`VK_USE_PLATFORM_IOS_MVK`

`VK_MVK_macos_surface`

macOS

`vulkan_macos.h`

None

`VK_USE_PLATFORM_MACOS_MVK`

`VK_OHOS_surface`

OHOS

`vulkan_ohos.h`

None

`VK_USE_PLATFORM_OHOS`

`VK_NN_vi_surface`

VI

`vulkan_vi.h`

None

`VK_USE_PLATFORM_VI_NN`

`VK_FUCHSIA_imagepipe_surface`

Fuchsia

`vulkan_fuchsia.h`

`<zircon/types.h>`

`VK_USE_PLATFORM_FUCHSIA`

`VK_EXT_metal_surface`

Metal on CoreAnimation

`vulkan_metal.h`

None

`VK_USE_PLATFORM_METAL_EXT`

`VK_QNX_screen_surface`

QNX Screen

`vulkan_screen.h`

`<screen/screen.h>`

`VK_USE_PLATFORM_SCREEN_QNX`

Note

This section describes the purpose of the headers independently of the specific underlying functionality of the window system extensions themselves. Each extension name will only link to a description of that extension when viewing a specification built with that extension included.

## [](#_see_also)See Also

[provisional-headers](provisional-headers.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-wsi-header)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0