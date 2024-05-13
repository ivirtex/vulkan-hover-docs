# LaunchSizeKHR(3) Manual Page

## Name

LaunchSizeKHR - Launch dimensions for ray shaders



## <a href="#_description" class="anchor"></a>Description

`LaunchSizeKHR`  
A variable decorated with the `LaunchSizeKHR` decoration will contain
the `width`, `height`, and `depth` dimensions passed to the
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html) command that initiated this
shader execution. The `width` is in the first component, the `height` is
in the second component, and the `depth` is in the third component.

Valid Usage

- <a href="#VUID-LaunchSizeKHR-LaunchSizeKHR-04269"
  id="VUID-LaunchSizeKHR-LaunchSizeKHR-04269"></a>
  VUID-LaunchSizeKHR-LaunchSizeKHR-04269  
  The `LaunchSizeKHR` decoration **must** be used only within the
  `RayGenerationKHR`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`,
  `MissKHR`, or `CallableKHR` `Execution` `Model`

- <a href="#VUID-LaunchSizeKHR-LaunchSizeKHR-04270"
  id="VUID-LaunchSizeKHR-LaunchSizeKHR-04270"></a>
  VUID-LaunchSizeKHR-LaunchSizeKHR-04270  
  The variable decorated with `LaunchSizeKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-LaunchSizeKHR-LaunchSizeKHR-04271"
  id="VUID-LaunchSizeKHR-LaunchSizeKHR-04271"></a>
  VUID-LaunchSizeKHR-LaunchSizeKHR-04271  
  The variable decorated with `LaunchSizeKHR` **must** be declared as a
  three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#LaunchSizeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
