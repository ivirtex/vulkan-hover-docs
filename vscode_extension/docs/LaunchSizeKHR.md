# LaunchSizeKHR(3) Manual Page

## Name

LaunchSizeKHR - Launch dimensions for ray shaders



## [](#_description)Description

`LaunchSizeKHR`

A variable decorated with the `LaunchSizeKHR` decoration will contain the `width`, `height`, and `depth` dimensions passed to the [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html) or [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysNV.html) command that initiated this shader execution. The `width` is in the first component, the `height` is in the second component, and the `depth` is in the third component.

Valid Usage

- [](#VUID-LaunchSizeKHR-LaunchSizeKHR-04269)VUID-LaunchSizeKHR-LaunchSizeKHR-04269  
  The `LaunchSizeKHR` decoration **must** be used only within the `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, `MissKHR`, or `CallableKHR` `Execution` `Model`
- [](#VUID-LaunchSizeKHR-LaunchSizeKHR-04270)VUID-LaunchSizeKHR-LaunchSizeKHR-04270  
  The variable decorated with `LaunchSizeKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-LaunchSizeKHR-LaunchSizeKHR-04271)VUID-LaunchSizeKHR-LaunchSizeKHR-04271  
  The variable decorated with `LaunchSizeKHR` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#LaunchSizeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0