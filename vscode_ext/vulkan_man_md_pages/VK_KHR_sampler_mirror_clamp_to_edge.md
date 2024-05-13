# VK_KHR_sampler_mirror_clamp_to_edge(3) Manual Page

## Name

VK_KHR_sampler_mirror_clamp_to_edge - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

15

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_sampler_mirror_clamp_to_edge%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_sampler_mirror_clamp_to_edge%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-08-17

**Contributors**  
- Tobias Hector, Imagination Technologies

- Jon Leech, Khronos

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_sampler_mirror_clamp_to_edge` extends the set of sampler address
modes to include an additional mode
(`VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`) that effectively uses a
texture map twice as large as the original image in which the additional
half of the new image is a mirror image of the original image.

This new mode relaxes the need to generate images whose opposite edges
match by using the original image to generate a matching “mirror image”.
This mode allows the texture to be mirrored only once in the negative s,
t, and r directions.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2.
However, if Vulkan 1.2 is supported and this extension is not, the
[VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html)
`VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` is optional. Since the
original extension did not use an author suffix on the enum
`VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`, it is used by both core
and extension implementations.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME`

- `VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_SPEC_VERSION`

- Extending [VkSamplerAddressMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html):

  - `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE`

  - `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR`

## <a href="#_example" class="anchor"></a>Example

Creating a sampler with the new address mode in each dimension

``` c
    VkSamplerCreateInfo createInfo =
    {
        .sType = VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO,
        // Other members set to application-desired values
    };

    createInfo.addressModeU = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE;
    createInfo.addressModeV = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE;
    createInfo.addressModeW = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE;

    VkSampler sampler;
    VkResult result = vkCreateSampler(
        device,
        &createInfo,
        &sampler);
```

## <a href="#_issues" class="anchor"></a>Issues

1\) Why are both KHR and core versions of the
`VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` token present?

**RESOLVED**: This functionality was intended to be required in Vulkan
1.0. We realized shortly before public release that not all
implementations could support it, and moved the functionality into an
optional extension, but did not apply the KHR extension suffix. Adding a
KHR-suffixed alias of the non-suffixed enum has been done to comply with
our own naming rules.

In a related change, before spec revision 1.1.121 this extension was
hardwiring into the spec Makefile so it was always included with the
Specification, even in the core-only versions. This has now been
reverted, and it is treated as any other extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-02-16 (Tobias Hector)

  - Initial draft

- Revision 2, 2019-08-14 (Jon Leech)

  - Add KHR-suffixed alias of non-suffixed enum.

- Revision 3, 2019-08-17 (Jon Leech)

  - Add an issue explaining the reason for the extension API not being
    suffixed with KHR.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_sampler_mirror_clamp_to_edge"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
